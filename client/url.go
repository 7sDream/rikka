package client

import (
	"encoding/json"
	"net/http"

	"github.com/7sDream/rikka/api"
)

func GetURL(host string, taskID string) (*api.URL, error) {
	url := host + api.URLPath + taskID
	l.Debug("Build url request url:", url)

	res, err := http.Get(url)
	if err != nil {
		l.Debug("Error happened when send url get request:", err)
		return nil, err
	}
	l.Debug("Send upload request successfully")

	resContent, err := checkRes(url, res)
	if err != nil {
		return nil, err
	}

	pURL := &api.URL{}

	if err := json.Unmarshal(resContent, pURL); err != nil || pURL.URL == "" {
		l.Debug("Decode response to url json failed, try to decode to error message")
		return nil, mustBeErrorJSON(resContent)
	}
	l.Debug("Decode response to url json successfully")
	return pURL, nil
}
