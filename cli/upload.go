package main

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/7sDream/rikka/api"
)

const uploadAPIPath = "/api/upload"

func createUploadRequest(url string, params map[string]string, paramName string, path string, content []byte) *http.Request {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		l.Fatal("Error happened when create form file:", err)
	}

	l.Debug("Create form writer successfully")

	if _, err = part.Write(content); err != nil {
		l.Fatal("Error happened when write file content to form file")
	}

	l.Debug("Write file content to form file successfully")

	for key, val := range params {
		if err = writer.WriteField(key, val); err != nil {
			l.Fatal("Error happened when write params[", key, "=", val, "] to form:", err)
		}
		l.Debug("Write params [", key, "=", val, "] to form successfully")
	}

	if err = writer.Close(); err != nil {
		l.Fatal("Error happened when close form writer:", err)
	}

	l.Debug("Close form writer successfully")

	req, err := http.NewRequest("POST", url, body)

	if err != nil {
		l.Fatal("Error happened when create request:", err)
	}

	l.Debug("Create request successfully")

	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req
}

func upload(host string, path string, content []byte, params map[string]string) string {
	client := &http.Client{}

	url := host + uploadAPIPath

	l.Debug("Build upload url:", url)

	req := createUploadRequest(url, params, uploadFileKey, path, content)

	res, err := client.Do(req)

	if err != nil {
		l.Fatal("Error when send upload request to", url, ":", err)
	}
	l.Debug("Send upload request successfully")

	resContent := checkRes(url, res)

	pTaskID := &api.TaskID{}

	if err := json.Unmarshal(resContent, pTaskID); err == nil {
		if pTaskID.TaskID != "" {
			l.Debug("Decode response to taskID json")
			return pTaskID.TaskID
		}
	}
	l.Debug("Decode response to taskID json failed, try to decode to error message")

	mustBeErrorJSON(resContent)

	return ""
}
