package weibo

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/textproto"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/7sDream/rikka/plugins"
)

var (
	weiboURL, _ = url.Parse("http://weibo.com")

	fileFieldKey   = "pic1"
	imageIDKey     = "pid"
	imageURLPrefix = "http://ww1.sinaimg.cn/large/"
	cbBase         = "http://weibo.com/aj/static/upimgback.html?_wv=5&callback=STK_ijax_"
	uploadURLBase  = "http://picupload.service.weibo.com/interface/pic_upload.php"
	uploadQuery    = struct {
		sync.Mutex
		M map[string]string
	}{
		M: map[string]string{
			"cb":      "",
			"url":     "x",
			"markpos": "1",
			"logo":    "0",
			"nick":    "x",
			"mask":    "0",
			"app":     "miniblog",
			"s":       "rdxt",
		},
	}

	queryLock = sync.Mutex{}
)

const (
	miniPublishPageURL = "http://weibo.com/minipublish"
)

func newWeiboClient() *http.Client {
	l.Debug("Creating weibo client")

	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		l.Fatal("Error happened when create cookiejar:", err)
	}
	l.Debug("Create cookiejar successfully")

	l.Debug("Create weibo client successfully")
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar: cookieJar,
	}
}

func updateCookies(cookieStr string) error {
	l.Debug("Updating cookies")
	rawRequest := fmt.Sprintf("GET / HTTP/1.0\r\nCookie: %s\r\n\r\n", cookieStr)

	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(rawRequest)))
	if err != nil {
		l.Error("Error when parse cookies from string", cookieStr, ":", err)
		return err
	}

	cookies := req.Cookies()

	if len(cookies) == 0 {
		errorMsg := "No cookies data in string your provided"
		l.Error(errorMsg)
		return errors.New(errorMsg)
	}

	l.Debug("Update cookies from string", cookieStr, "successfully")
	for _, cookie := range cookies {
		l.Debug(fmt.Sprintf("%#v", cookie))
	}

	client.Jar.SetCookies(weiboURL, cookies)
	return nil
}

func auxCheckLogin() (bool, error) {
	l.Debug("Checking is login...")
	res, err := client.Get(miniPublishPageURL)
	if err != nil {
		l.Error("Error happened when visit minipublish page:", err)
		return false, err
	}
	l.Debug("Visit minipublish page successfully, code", res.StatusCode)
	defer res.Body.Close()
	return res.StatusCode == http.StatusOK, nil
}

func auxUpdateCB() {
	uploadQuery.M["cb"] = cbBase + strconv.FormatInt(time.Now().Unix(), 10)
}

func auxGetUploadURL() string {
	uploadQuery.Lock()
	defer uploadQuery.Unlock()

	auxUpdateCB()

	uploadURL, _ := url.Parse(uploadURLBase)
	query := uploadURL.Query()
	for key, val := range uploadQuery.M {
		query.Set(key, val)
	}
	uploadURL.RawQuery = query.Encode()

	return uploadURL.String()
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func auxCreateImageFormFileField(w *multipart.Writer, fileFieldKey string, filename string, fileType string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fileFieldKey), escapeQuotes(filename)))
	h.Set("Content-Type", "image/"+fileType)
	return w.CreatePart(h)
}

func auxCreateUploadRequest(q *plugins.SaveRequest) (*http.Request, error) {
	l.Debug("Creating upload request...")

	auxUpdateCB()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := auxCreateImageFormFileField(writer, fileFieldKey, "noname."+q.FileExt, q.FileExt)
	if err != nil {
		l.Error("Error happened when create form file:", err)
		return nil, err
	}
	l.Debug("Create form writer successfully")

	content, err := ioutil.ReadAll(q.File)
	if err != nil {
		l.Error("Error happened when read file content")
		return nil, err
	}
	l.Debug("Read file content successfully")

	if _, err = part.Write(content); err != nil {
		l.Error("Error happened when write file content to form:", err)
		return nil, err
	}
	l.Debug("Write file content to form file successfully")

	if err = writer.Close(); err != nil {
		l.Error("Error happened when close form writer:", err)
		return nil, err
	}
	l.Debug("Close form writer successfully")

	uploadURL := auxGetUploadURL()
	req, err := http.NewRequest("POST", uploadURL, body)
	if err != nil {
		l.Error("Error happened when create post request with url", uploadURL, ":", err)
		return nil, err
	}

	var cookieStrs []string
	for _, cookie := range client.Jar.Cookies(weiboURL) {
		cookieStrs = append(cookieStrs, cookie.Name+"="+cookie.Value)
	}

	req.Header.Set("Cookie", strings.Join(cookieStrs, "; "))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	l.Debug("Create post request successfully, url", uploadURL)

	return req, nil
}

func auxGetImageURL(raw string) (string, error) {
	l.Debug("Getting image url from redirect url", raw)

	redirectURL, err := url.Parse(raw)
	if err != nil {
		l.Error("Error happened when parse redirect URL ", redirectURL, ":", err)
		return "", err
	}
	l.Debug("parse redirect URL", raw, "successfully")

	imageID := redirectURL.Query().Get(imageIDKey)
	if imageID == "" {
		errorMsg := "No image ID field " + imageIDKey + " in url " + raw + ", weibo api changed"
		l.Error("Error happened when get image id:", errorMsg)
		return "", errors.New(errorMsg)
	}
	l.Debug("Get image ID", imageID, "from url", raw, "successfully")

	return imageURLPrefix + imageID, nil
}

func auxUpload(q *plugins.SaveRequest) (string, error) {
	l.Debug("Truely uploading image...")

	req, err := auxCreateUploadRequest(q)
	if err != nil {
		l.Error("Error happened when create upload request:", err)
		return "", err
	}
	l.Debug("Create upload request successfully")

	res, err := client.Do(req)
	if err != nil {
		l.Error("Error happened when send upload request:", err)
		return "", err
	}

	if res.StatusCode != http.StatusFound {
		errorMsg := "Upload response code is not 302, weibo api changed"
		l.Error("Error happened when get image url:", errorMsg)
		return "", errors.New(errorMsg)
	}

	redirectURL := res.Header.Get("Location")
	if redirectURL == "" {
		errorMsg := "No location header, weibo api changed"
		l.Error("Error happened when get image url:", errorMsg)
		return "", errors.New(errorMsg)
	}

	imageURL, err := auxGetImageURL(redirectURL)
	if err != nil {
		errorMsg := "can't get image url from Location header, weibo api changed"
		l.Error("Error happened when get image:", errorMsg)
		return "", errors.New(errorMsg)
	}

	return imageURL, nil
}

func upload(client *http.Client, q *plugins.SaveRequest) (string, error) {
	l.Debug("Preparing upload...")

	login, err := auxCheckLogin()
	if err != nil {
		l.Error("Error happened when check login:", err)
		return "", err
	}
	l.Debug("Check login successfully")

	if !login {
		l.Error("No weibo account login")
		return "", errors.New("Weibo account not login, please set cookies")
	}
	l.Debug("Weibo account is logged")

	imageURL, err := auxUpload(q)
	if err != nil {
		l.Error("Error happened when get image url:", err)
		return "", err
	}
	l.Debug("Get image url", imageURL, "successfully")

	return imageURL, nil
}
