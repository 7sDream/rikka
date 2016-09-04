package main

import (
	"encoding/json"
	"net/http"

	"github.com/7sDream/rikka/api"
)

func getURL(host string, taskID string) *api.URL {

	url := host + api.URLPath + taskID
	l.Debug("Build url request url:", url)

	res, err := http.Get(url)
	if err != nil {
		l.Fatal("Error when send url request to", url, ":", err)
	}
	l.Debug("Send upload request successfully")

	resContent := checkRes(url, res)

	pURL := &api.URL{}

	if err := json.Unmarshal(resContent, pURL); err == nil {
		if pURL.URL != "" {
			l.Debug("Decode response to url json")
			return pURL
		}
	}
	l.Debug("Decode response to url json failed, try to decode to error message")

	mustBeErrorJSON(resContent)

	return nil
}
