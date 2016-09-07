package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	pathutil "path/filepath"
	"strings"

	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/server/apiserver"
)

func getFile() (string, []byte) {
	if len(flag.Args()) != 1 {
		l.Fatal("No file specified or more than one file specified")
	}
	filepath := flag.Args()[0]
	l.Debug("Get path of file want be uploaded:", filepath)

	absFilePath, err := pathutil.Abs(filepath)
	if err != nil {
		l.Fatal(filepath, "is not a file path")
	}
	l.Debug("Change to absolute path:", absFilePath)

	if !util.IsDir(absFilePath) {
		l.Fatal("Path ", absFilePath, "not exists or not a dir")
	}
	l.Debug("File", absFilePath, "exists and is a file")

	fileContent, err := ioutil.ReadFile(absFilePath)
	if err != nil {
		l.Fatal("Error happened when read file", filepath, ":", err)
	}
	l.Info("Read file", absFilePath, "content successfully")

	filetype := http.DetectContentType(fileContent)
	if !apiserver.IsAccepted(filetype) {
		l.Fatal("File", absFilePath, "is not a image file, it is", filetype)
	}
	l.Debug("Fie", absFilePath, "type check passed:", filetype)

	return absFilePath, fileContent
}
