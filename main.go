package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"

	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/plugins"
	"github.com/7sDream/rikka/plugins/fs"
	"github.com/7sDream/rikka/server"
)

// plugin name to Plugin object map
var pluginMap = make(map[string]plugins.RikkaPlugin)

// command line arguments
var argBindIPAddress *string
var argPort *int
var argPassword *string
var argMaxSizeByMB *float64
var argPluginStr *string

// the plugin
var thePlugin plugins.RikkaPlugin

// logger
var l = logger.NewLogger("[Main]")

// --- init functions ---

func init() {
	initPluginList()

	initArgVars()
	flag.Parse()
	l.Info("Args bindIP =", *argBindIPAddress)
	l.Info("Args port =", *argPort)
	l.Info("Args password =", *argPassword)
	l.Info("Args maxFileSize =", *argMaxSizeByMB, "MB")
	l.Info("Args.plugin =", *argPluginStr)

	runtimeCheck()
}

func initPluginList() {
	pluginMap["fs"] = fs.FsPlugin
}

func initArgVars() {
	argBindIPAddress = flag.String("bind", ":", "bind ip address, use : for all address")
	argPort = flag.Int("port", 80, "server port")
	argPassword = flag.String("pwd", "rikka", "The password need provided when upload")
	argMaxSizeByMB = flag.Float64("size", 5, "Max file size by MB")

	pluginNames := make([]string, 0, len(pluginMap))
	for k := range pluginMap {
		pluginNames = append(pluginNames, k)
	}
	argPluginStr = flag.String(
		"plugin", "fs",
		"what plugin use to save file, selected from "+fmt.Sprintf("%v", pluginNames),
	)

}

func runtimeCheck() {
	l.Info("Check runtime environment")

	l.Info("Try to find plugin", *argPluginStr)
	if plugin, ok := pluginMap[*argPluginStr]; ok {
		thePlugin = plugin
		l.Info("Plugin", *argPluginStr, "found")
	} else {
		l.Fatal("Plugin", *argPluginStr, "not exist")
	}

	l.Info("All runtime environment check passed")
}

// main enterypoint

func SignalHandler(c chan os.Signal) func() {
	return func() {
		for _ = range c {
			l.Info("Rikka need go to sleep, see you tomorrow")
			os.Exit(0)
		}
	}
}

func main() {

	// handler Ctrl + C
	signalChain := make(chan os.Signal, 1)
	signal.Notify(signalChain, os.Interrupt)
	go SignalHandler(signalChain)()

	var socket string

	if *argBindIPAddress == ":" {
		socket = *argBindIPAddress + strconv.Itoa(*argPort)
	} else {
		socket = *argBindIPAddress + ":" + strconv.Itoa(*argPort)
	}

	l.Info(
		"Try start rikka at", socket,
		", with password", *argPassword,
		", max file size", *argMaxSizeByMB, "MB",
		", plugin", *argPluginStr,
	)

	server.StartRikkaServer(socket, *argPassword, *argMaxSizeByMB, thePlugin)
}
