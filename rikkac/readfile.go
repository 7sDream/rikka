package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	pathutil "path/filepath"
	"strings"

	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/server/apiserver"
)

func getFile() (string, []byte) {
	filepath := ""
	if len(flag.Args()) == 1 {
		filepath = flag.Args()[0]
	} else {
		if !strings.HasPrefix(os.Args[1], "-") {
			filepath = os.Args[1]
		} else {
			l.Fatal("No or more than one file specified")
		}
	}

	l.Debug("Get path of file want be uploaded:", filepath)

	absFilePath, err := pathutil.Abs(filepath)
	if err != nil {
		l.Fatal(filepath, "is not a file path")
	}
	l.Debug("Change to absolute path:", absFilePath)

	if !util.IsFile(absFilePath) {
		l.Fatal("Path ", absFilePath, "not exists or not a file")
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
