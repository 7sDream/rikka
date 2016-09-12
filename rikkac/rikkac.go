package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/7sDream/rikka/common/logger"
)

const (
	version = "0.0.5"
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

func waitOutput(index int, out chan *taskRes) {
	if index == 0 {
		l.Fatal("No file provided")
	} else if index == 1 {
		c := <-out
		fmt.Println(c.StringWithoutFilepath())
	} else {
		nowShow := 0
		resList := make([]*taskRes, index)
		for i := 0; i < index; i++ {
			c := <-out
			resList[c.Index] = c
			if c.Index == nowShow {
				for nowShow < index && resList[nowShow] != nil {
					c = resList[nowShow]
					fmt.Println(c.String())
					nowShow++
				}
			}
		}
	}
}

func main() {

	host := getHost()

	index := 0
	ok := true
	out := make(chan *taskRes)
	var filepath string

	for ok {
		filepath, ok = getFile(index)
		if ok {
			l.Info("Read image file", filepath, "successfully, add to task list")
			go worker(host, filepath, index, out)
			index++
		}
	}
	l.Info("End with index", index)

	waitOutput(index, out)
}
