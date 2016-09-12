package client

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/7sDream/rikka/api"
)

func createUploadRequest(url string, path string, content []byte, params map[string]string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(api.FormKeyFile, filepath.Base(path))
	if err != nil {
		l.Debug("Error happened when create form file:", err)
		return nil, err
	}
	l.Debug("Create form writer successfully")

	if _, err = part.Write(content); err != nil {
		l.Debug("Error happened when write file content to form:", err)
		return nil, err
	}
	l.Debug("Write file content to form file successfully")

	for key, val := range params {
		if err = writer.WriteField(key, val); err != nil {
			l.Debug("Error happened when try to write params [", key, "=", val, "] to form:", err)
			return nil, err
		}
		l.Debug("Write params [", key, "=", val, "] to form successfully")
	}

	if err = writer.Close(); err != nil {
		l.Debug("Error happened when close form writer:", err)
		return nil, err
	}
	l.Debug("Close form writer successfully")

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		l.Debug("Error happened when create post request:", err)
		return nil, err
	}
	l.Debug("Create request successfully")

	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}

func getParams(password string) map[string]string {
	params := map[string]string{
		api.FormKeyFrom: api.FromAPI,
		api.FormKeyPWD:  password,
	}

	l.Debug("Build params:", params)

	return params
}

func Upload(host string, path string, content []byte, password string) (string, error) {
	client := &http.Client{}

	url := host + api.UploadPath

	l.Debug("Build upload url:", url)

	req, err := createUploadRequest(url, path, content, getParams(password))
	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		l.Debug("Error happened when try to send upload request:", err)
		return "", err
	}
	l.Debug("Send upload request successfully")

	resContent, err := checkRes(url, res)
	if err != nil {
		return "", err
	}

	pTaskID := &api.TaskID{}
	if err = json.Unmarshal(resContent, pTaskID); err != nil || pTaskID.TaskID == "" {
		l.Debug("Decode response to taskID json failed, try to decode to error message")
		return "", mustBeErrorJSON(resContent)
	}
	l.Debug("Decode response to taskID json successfully")

	return pTaskID.TaskID, nil
}
