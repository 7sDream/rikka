package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/7sDream/rikka/common/logger"
)

const (
	version = "0.0.3"
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
	params := getParams()
	filePath, fileContent := getFile()

	l.Info("Start upload")

	taskID := upload(host, filePath, fileContent, params)
	l.Info("Get taskID:", taskID)

	waitFinish(host, taskID)
	l.Info("Task state comes to finished")

	pURL := getURL(host, taskID)
	l.Info("Url gotten:", *pURL)

	formatted := format(pURL)
	l.Info("Make final formatted text successfully:", formatted)

	fmt.Println(formatted)
}
