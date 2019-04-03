package cos

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"

	"github.com/7sDream/rikka/plugins"
)

const (
	uploadBaseURL     = "http://web.file.myqcloud.com/files/v1/%s/%s/%s%s"
	taskIDPlaceholder = "{taskID}"
	baseCosV5Url      = "http://%s-%s.cos.%s.myqcloud.com"
)

type cosClient struct {
	http.Client
	sign   string
	expire time.Time
}

type cosSdkV5Client struct {
	*cos.Client
}

type genericCosClient interface {
	Upload(q *plugins.SaveRequest, taskID string) error
}

func makeSign(current *time.Time, dur *time.Duration, randInt *int) (string, time.Time) {
	var t time.Time
	if current == nil {
		t = time.Now().UTC()
	} else {
		t = *current
	}
	if dur == nil {
		durObj := 90 * 24 * time.Hour
		dur = &durObj
	}
	e := t.Add(*dur)
	if randInt == nil {
		number := rand.Intn(10000000000)
		randInt = &number
	}
	// Original = "a=[appid]&b=[bucket]&k=[SecretID]&e=[expiredTime]&t=[currentTime]&r=[rand]&f="
	original := fmt.Sprintf(
		"a=%s&b=%s&k=%s&e=%d&t=%d&r=%d&f=",
		appID, bucketName, secretID, e.Unix(), // 60 * 60 * 24 * 90 = 90 days
		t.Unix(), *randInt, // random integer, max length: 10
	)
	hmacBuilder := hmac.New(sha1.New, []byte(secretKey))
	hmacBuilder.Write([]byte(original))
	signTemp := hmacBuilder.Sum(nil)
	sign := base64.StdEncoding.EncodeToString(append(signTemp, []byte(original)...))
	return sign, e
}

func newCosClient() *cosClient {
	sign, expire := makeSign(nil, nil, nil)
	return &cosClient{
		Client: http.Client{},
		sign:   sign,
		expire: expire,
	}
}

func newCosSdkV5Client() *cosSdkV5Client {
	rawUrl := fmt.Sprintf(baseCosV5Url, bucketName, appID, region)
	u, _ := url.Parse(rawUrl)
	return &cosSdkV5Client{
		cos.NewClient(&cos.BaseURL{BucketURL: u}, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  secretID,
				SecretKey: secretKey,
			},
		}),
	}
}

func (c *cosClient) auxMakeUploadRequest(q *plugins.SaveRequest, taskID string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("fileContent", taskID)
	if err != nil {
		l.Error("Error happened when create form file of task", taskID, ":", err)
		return nil, err
	}
	l.Debug("Create form writer of task", taskID, "successfully")

	fileContent, err := ioutil.ReadAll(q.File)

	//noinspection GoUnhandledErrorResult
	defer q.File.Close()

	if err != nil {
		l.Error("Error happened when read file content of task", taskID, ":", err)
		return nil, err
	}
	l.Debug("Read file content of task", taskID, "successfully")

	if _, err = part.Write(fileContent); err != nil {
		l.Debug("Error happened when write file content of task", taskID, "to form:", err)
		return nil, err
	}
	l.Debug("Write file content of task", taskID, "to form file successfully")

	shaOfFile := sha1.Sum(fileContent)
	shaString := strings.ToUpper(hex.EncodeToString(shaOfFile[:]))
	l.Info("Get sha256 of task", taskID, ":", shaString)

	params := map[string]string{
		"op":         "upload",
		"sha":        shaString,
		"insertOnly": "0",
	}

	for key, val := range params {
		if err = writer.WriteField(key, val); err != nil {
			l.Error("Error happened when try to write params [", key, "=", val, "] to form in task", taskID, ":", err)
			return nil, err
		}
		l.Debug("Write params [", key, "=", val, "] to form in task", taskID, "successfully")
	}

	if err = writer.Close(); err != nil {
		l.Debug("Error happened when close form writer of task", taskID, ":", err)
		return nil, err
	}
	l.Debug("Close form writer of task", taskID, "successfully")

	rawUrl := fmt.Sprintf(uploadBaseURL, appID, bucketName, bucketPath, taskID)

	req, err := http.NewRequest("POST", rawUrl, body)
	if err != nil {
		l.Debug("Error happened when create post request of task", taskID, ":", err)
		return nil, err
	}
	l.Debug("Create request of task", taskID, "successfully")

	req.Header.Set("Content-Type", writer.FormDataContentType())

	if time.Now().UTC().Add(time.Hour).After(c.expire) {
		newSign, newExpire := makeSign(nil, nil, nil)
		c.sign = newSign
		c.expire = newExpire
		l.Info("Renew sign, next expire date:", newExpire)
	}

	req.Header.Set("Authorization", c.sign)

	return req, nil
}

func (c *cosClient) Upload(q *plugins.SaveRequest, taskID string) error {
	req, err := c.auxMakeUploadRequest(q, taskID)
	if err != nil {
		l.Error("Error happened when create upload request of task", taskID, ":", err)
		return err
	}
	l.Debug("Create upload request of task", taskID, "successfully")

	res, err := c.Do(req)
	if err != nil {
		l.Error("Error happened when send request or get response of task", taskID, ":", err)
		return err
	}
	l.Debug("Send request and get response of task", taskID, "successfully")

	resContent, err := ioutil.ReadAll(res.Body)

	//noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	if err != nil {
		l.Error("Error when read response body of task", taskID, ":", err)
		return err
	}
	l.Debug("Read response body of task", taskID, "successfully")

	var resJSON interface{}
	err = json.Unmarshal(resContent, &resJSON)
	if err != nil {
		l.Error("Error happened when parer response body as json:", err)
		return err
	}

	m := resJSON.(map[string]interface{})
	jsonString := fmt.Sprintf("%#v", m)

	l.Info("Get response json:", jsonString)

	code := m["code"].(float64)
	if code != 0 {
		errorMsg := m["message"].(string)
		l.Error("Error happened when upload", taskID, ":", errorMsg)
		return errors.New(errorMsg)
	}
	l.Debug("Image upload of task", taskID, "successfully")

	if bucketHost == "" {
		data := m["data"].(map[string]interface{})
		accessUrl := data["access_url"].(string)
		bucketHost = strings.Replace(accessUrl, taskID, taskIDPlaceholder, -1)
		l.Debug("Get image accessUrl format:", bucketHost)
	}

	return nil
}

func buildPath(taskID string) string {
	if len(bucketPath) > 0 {
		return bucketPath + "/" + taskID
	}
	return taskID
}

func (c *cosSdkV5Client) Upload(q *plugins.SaveRequest, taskID string) error {
	_, e := c.Client.Object.Put(context.Background(), buildPath(taskID), q.File, nil)
	bucketHost = strings.Replace(fmt.Sprintf(baseCosV5Url+"/%s%s", bucketName, appID, region, bucketPath, taskID), taskID, taskIDPlaceholder, -1)
	return e
}
