package main

import (
	"flag"
	"os"
	pathutil "path/filepath"
	"strings"

	"github.com/7sDream/rikka/client"
	"github.com/7sDream/rikka/common/util"
)

func getFile() (string, []byte, error) {
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

	fileContent, err := client.CheckFile(absFilePath)
	if err != nil {
		return "", nil, err
	}
	return absFilePath, fileContent, nil
}
