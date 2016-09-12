package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/7sDream/rikka/client"
	"github.com/7sDream/rikka/common/logger"
)

const (
	version = "0.0.4"
)

var (
	l = logger.NewLogger("[CLI]")

	argInfo    = flag.Bool("v", false, "set logger level to Info")
	argDebug   = flag.Bool("vv", false, "set logger level to Debug")
	argVersion = flag.Bool("version", false, "show rikkac version and exit")
)

func init() {
	flag.Parse()

	if *argVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if *argDebug {
		logger.SetLevel(logger.LevelDebug)
	} else if *argInfo {
		logger.SetLevel(logger.LevelInfo)
	} else {
		logger.SetLevel(logger.LevelWarn)
	}
}

func main() {

	host := getHost()
	filePath, fileContent, err := getFile()
	if err != nil {
		l.Fatal("Error happened when try to read image file:", err)
	}
	l.Info("Read image file successfully")

	taskID, err := client.Upload(host, filePath, fileContent, getPassword())
	if err != nil {
		l.Fatal("Error happened when upload image:", err)
	}
	l.Info("Upload successfully, get taskID:", taskID)

	err = client.WaitFinish(host, taskID)
	if err != nil {
		l.Fatal("Error happened when wait state to finished:", err)
	}
	l.Info("Task state comes to finished")

	pURL, err := client.GetURL(host, taskID)
	if err != nil {
		l.Fatal("Error happened when get url of image:", err)
	}
	l.Info("Url gotten:", *pURL)

	formatted := format(pURL)
	l.Info("Make final formatted text successfully:", formatted)

	fmt.Println(formatted)
}
