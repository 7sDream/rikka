package client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/7sDream/rikka/api"
)

func mustBeErrorJSON(content []byte) error {
	pError := &api.Error{}
	var err error
	if err = json.Unmarshal(content, pError); err != nil {
		l.Debug("Error happened when decode response to error json:", err)
		return err
	}
	if pError.Error == "" {
		l.Debug("Unable to decode response to error json:", "result is empty")
		return errors.New("Unable to decode Rikka server response to error json" + string(content))
	}
	return errors.New(pError.Error)
}

func checkRes(url string, res *http.Response) ([]byte, error) {
	resContent, err := ioutil.ReadAll(res.Body)
	if err != nil {
		l.Debug("Error happaned when read response body:", err)
		return nil, err
	}
	l.Debug("Get response content of", url, "successfully:", string(resContent))

	if res.StatusCode != http.StatusOK {
		l.Debug("Rikka return a non-ok status code", res.StatusCode)
		return nil, mustBeErrorJSON(resContent)
	}
	l.Debug("Rikka response OK when request", url)

	return resContent, nil
}
