package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/7sDream/rikka/api"
)

func mustBeErrorJSON(content []byte) {
	pError := &api.Error{}
	if err := json.Unmarshal(content, pError); err == nil {
		if pError.Error != "" {
			l.Debug("Decode response to error json")
			l.Fatal("Rikka server return a error message:", pError.Error)
		}
	}
	l.Fatal("Unable to decode Rikka server response", string(content))
}

func checkRes(url string, res *http.Response) []byte {
	resContent, err := ioutil.ReadAll(res.Body)
	if err != nil {
		l.Fatal("Error when get response content of", url, ":", err)
	}
	l.Debug("Get response content of", url, "successfully:", string(resContent))

	if res.StatusCode != http.StatusOK {
		l.Error("Rikka server return a non-ok statu code", res.StatusCode, "when request", url)
		mustBeErrorJSON(resContent)
	}
	l.Debug("Rikka response OK when request", url)

	return resContent
}
