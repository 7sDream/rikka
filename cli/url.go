package main

import (
	"encoding/json"
	"net/http"

	"github.com/7sDream/rikka/plugins"
)

const urlAPIPath = "/api/url/"

func getURL(host string, taskID string) *plugins.URLJSON {

	url := host + urlAPIPath + taskID
	l.Debug("Build url request url:", url)

	res, err := http.Get(url)
	if err != nil {
		l.Fatal("Error when send url request to", url, ":", err)
	}
	l.Debug("Send upload request successfully")

	resContent := checkRes(url, res)

	urlJSON := &plugins.URLJSON{}

	if err := json.Unmarshal(resContent, &urlJSON); err == nil {
		if urlJSON.URL != "" {
			l.Debug("Decode response to url json")
			return urlJSON
		}
	}
	l.Debug("Decode response to url json failed, try to decode to error message")

	mustBeErrorJSON(resContent)

	return nil
}
