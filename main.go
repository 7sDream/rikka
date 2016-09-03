package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
	"github.com/7sDream/rikka/plugins/fs"
)

var pluginMap = make(map[string]plugins.RikkaPlugin)

var argPluginStr *string
var argPort *int
var argPassword *string
var argMaxSizeByMB *float64

var l = logger.NewLogger("[Main]")

func initPluginList() {
	pluginMap["fs"] = fs.FsPlugin
}

func initArgVars() {
	pluginNames := make([]string, 0, len(pluginMap))
	for k := range pluginMap {
		pluginNames = append(pluginNames, k)
	}

	argPluginStr = flag.String(
		"plugin", "fs",
		"what plugin use to save file, selected from "+fmt.Sprintf("%v", pluginNames),
	)

	argPort = flag.Int("port", 80, "server port")
	argPassword = flag.String("pwd", "rikka", "The password need provided when upload")
	argMaxSizeByMB = flag.Float64("size", 5, "Max file size by MB")
}

func runtimeCheck() {
	l.Info("Check runtime environment")

	requireFiles := []string{
		"templates", "templates/index.html", "templates/view.html",
		"static", "static/main.css", "static/index.css", "static/view.css", "static/rikka.png",
	}

	for _, file := range requireFiles {
		if util.CheckExist(file) {
			l.Info("Needed", file, "exist, check passed")
		} else {
			l.Fatal(file, "not exist, check failed, exit")
		}
	}

	l.Info("Try to find plugin", *argPluginStr)
	if _, ok := pluginMap[*argPluginStr]; ok {
		l.Info("Plugin", *argPluginStr, "found")
	} else {
		l.Fatal("Plugin", *argPluginStr, "not exist")
	}

	l.Info("All runtime environment check passed")
}

func init() {
	initPluginList()

	initArgVars()
	flag.Parse()
	l.Info("Args port =", *argPort)
	l.Info("Args password =", *argPassword)
	l.Info("Args maxFileSize =", *argMaxSizeByMB, "MB")
	l.Info("Args.plugin =", *argPluginStr)

	runtimeCheck()
}

func main() {
	staticFs := util.DisableListDir(http.FileServer(http.Dir("static")))

	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/view/", view)
	http.Handle("/static/", http.StripPrefix("/static", staticFs))

	l.Info("Rikka main part started successfully")
	l.Info("Loading plugin", *argPluginStr)
	plugins.Load(pluginMap[*argPluginStr])

	err := http.ListenAndServe(":"+strconv.Itoa(*argPort), nil)
	if err != nil {
		l.Fatal(err.Error())
	}
}
