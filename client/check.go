package client

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/7sDream/rikka/server/apiserver"
)

func CheckFile(absFilePath string) ([]byte, error) {
	fileContent, err := ioutil.ReadFile(absFilePath)
	if err != nil {
		l.Debug("Error happened when try to read file", absFilePath, ":", err)
		return nil, err
	}
	l.Debug("Read file", absFilePath, "content successfully")

	filetype := http.DetectContentType(fileContent)
	if _, ok := apiserver.IsAccepted(filetype); !ok {
		errMsg := "File" + absFilePath + "is not a acceptable image file, it is" + filetype
		l.Debug(errMsg)
		return nil, errors.New(errMsg)
	}
	l.Debug("Fie", absFilePath, "type check passed:", filetype)

	return fileContent, nil
}
