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

	fileType := http.DetectContentType(fileContent)
	if _, ok := apiserver.IsAccepted(fileType); !ok {
		errMsg := "File" + absFilePath + "is not a acceptable image file, it is" + fileType
		l.Debug(errMsg)
		return nil, errors.New(errMsg)
	}
	l.Debug("Fie", absFilePath, "type check passed:", fileType)

	return fileContent, nil
}
