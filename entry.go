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

// Map from plugin name to object
var pluginMap = make(map[string]plugins.RikkaPlugin)

// Command line arguments var
var argBindIPAddress *string
var argPort *int
var argPassword *string
var argMaxSizeByMB *float64
var argPluginStr *string
var argLogLevel *int

// The used plugin
var thePlugin plugins.RikkaPlugin

// Logger of this package
var l = logger.NewLogger("[Entry]")

// --- Init and check functions ---

func init() {
	initPluginList()

	initArgVars()
	flag.Parse()
	l.Info("Args bindIP =", *argBindIPAddress)
	l.Info("Args port =", *argPort)
	l.Info("Args password =", *argPassword)
	l.Info("Args maxFileSize =", *argMaxSizeByMB, "MB")
	l.Info("Args loggerLevel =", *argLogLevel)
	l.Info("Args.plugin =", *argPluginStr)

	logger.SetLevel(*argLogLevel)
	runtimeEnvCheck()
}

func initPluginList() {
	pluginMap["fs"] = fs.FsPlugin
}

func initArgVars() {
	argBindIPAddress = flag.String("bind", ":", "bind ip address, use : for all address")
	argPort = flag.Int("port", 80, "server port")
	argPassword = flag.String("pwd", "rikka", "The password need provided when upload")
	argMaxSizeByMB = flag.Float64("size", 5, "Max file size by MB")
	argLogLevel = flag.Int(
		"level", logger.LevelInfo,
		fmt.Sprintf("logger level, from %d to %d", logger.LevelInfo, logger.LevelError),
	)

	// Get name array of all avaliable plugins, show in `rikka -h``
	pluginNames := make([]string, 0, len(pluginMap))
	for k := range pluginMap {
		pluginNames = append(pluginNames, k)
	}
	argPluginStr = flag.String(
		"plugin", "fs",
		"what plugin use to save file, selected from "+fmt.Sprintf("%v", pluginNames),
	)

}

func runtimeEnvCheck() {
	l.Info("Check runtime environment")

	l.Info("Try to find plugin", *argPluginStr)

	// Make sure plugin be selected exist
	if plugin, ok := pluginMap[*argPluginStr]; ok {
		thePlugin = plugin
		l.Info("Plugin", *argPluginStr, "found")
	} else {
		l.Fatal("Plugin", *argPluginStr, "not exist")
	}

	l.Info("All runtime environment check passed")
}

func createSignalHandler(c chan os.Signal) func() {
	return func() {
		for _ = range c {
			l.Info("Rikka have to go to bed, see you tomorrow")
			os.Exit(0)
		}
	}
}

// Main enterypoint

func main() {

	// handler Ctrl + C
	signalChain := make(chan os.Signal, 1)
	signal.Notify(signalChain, os.Interrupt)
	signalHandler := createSignalHandler(signalChain)
	go signalHandler()

	// concat socket from ip address and port
	var socket string

	if *argBindIPAddress == ":" {
		socket = *argBindIPAddress + strconv.Itoa(*argPort)
	} else {
		socket = *argBindIPAddress + ":" + strconv.Itoa(*argPort)
	}

	// print launch args
	l.Info(
		"Try start rikka at", socket,
		", with password", *argPassword,
		", max file size", *argMaxSizeByMB, "MB",
		", plugin", *argPluginStr,
	)

	// start Rikka server (this call is Sync)
	server.StartRikka(socket, *argPassword, *argMaxSizeByMB, thePlugin)
}
