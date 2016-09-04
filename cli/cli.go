package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/7sDream/rikka/common/logger"
)

var l = logger.NewLogger("[CLI]")

var argInfo = flag.Bool("v", false, "set logger level to Info")
var argDebug = flag.Bool("vv", false, "set logger level to Debug")

func init() {
	flag.Parse()
	if *argDebug {
		logger.SetLevel(logger.LevelDebug)
	} else if *argInfo {
		logger.SetLevel(logger.LevelInfo)
	} else {
		logger.SetLevel(logger.LevelWarn)
	}
}

func isDir(filepath string) bool {
	stat, _ := os.Stat(filepath)
	return stat.IsDir()
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

	url := getURL(host, taskID)

	l.Info("Url gotten:", url)

	formatted := format(url)

	l.Info("Make final formatted text successfully:", formatted)

	fmt.Println(formatted)
}
