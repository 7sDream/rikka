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
	"github.com/7sDream/rikka/plugins/qiniu"
	"github.com/7sDream/rikka/plugins/tencent/ci"
	"github.com/7sDream/rikka/plugins/tencent/cos"
	"github.com/7sDream/rikka/plugins/upai"
	"github.com/7sDream/rikka/plugins/weibo"
)

var (
	// Map from plugin name to object
	pluginMap = make(map[string]plugins.RikkaPlugin)

	// Command line arguments var
	argBindIPAddress *string
	argPort          *int
	argPassword      *string
	argMaxSizeByMB   *float64
	argPluginStr     *string
	argLogLevel      *int
	argHTTPS         *bool
	argCertDir       *string
	argAllowOrigin   *string

	// concat socket from ip address and port
	socket string

	// The used plugin
	thePlugin plugins.RikkaPlugin
)

// --- Init and check ---

func createSignalHandler(handlerFunc func()) (func(), chan os.Signal) {
	signalChain := make(chan os.Signal, 1)

	return func() {
		for range signalChain {
			handlerFunc()
		}
	}, signalChain
}

// registerSignalHandler register a handler for process Ctrl + C
func registerSignalHandler(handlerFunc func()) {
	signalHandler, channel := createSignalHandler(handlerFunc)
	signal.Notify(channel, os.Interrupt)
	go signalHandler()
}

func init() {

	registerSignalHandler(func() {
		l.Info("Receive interrupt signal")
		l.Info("Rikka have to go to sleep, see you tomorrow")
		os.Exit(0)
	})

	initPluginList()

	initArgVars()

	flag.Parse()

	l.Info("Args bindIP =", *argBindIPAddress)
	l.Info("Args port =", *argPort)
	l.Info("Args password =", *argPassword)
	l.Info("Args maxFileSize =", *argMaxSizeByMB, "MB")
	l.Info("Args loggerLevel =", *argLogLevel)
	l.Info("Args https =", *argHTTPS)
	l.Info("Args cert dir =", *argCertDir)
	l.Info("Args plugin =", *argPluginStr)

	if *argPort == 0 {
		if *argHTTPS {
			*argPort = 443
		} else {
			*argPort = 80
		}
	}

	if *argBindIPAddress == ":" {
		socket = *argBindIPAddress + strconv.Itoa(*argPort)
	} else {
		socket = *argBindIPAddress + ":" + strconv.Itoa(*argPort)
	}

	logger.SetLevel(*argLogLevel)

	runtimeEnvCheck()
}

func initPluginList() {
	pluginMap["fs"] = fs.Plugin
	pluginMap["qiniu"] = qiniu.Plugin
	pluginMap["upai"] = upai.Plugin
	pluginMap["weibo"] = weibo.Plugin
	pluginMap["tccos"] = cos.Plugin
	pluginMap["tcci"] = ci.Plugin
}

func initArgVars() {
	argBindIPAddress = flag.String("bind", ":", "Bind ip address, use : for all address")
	argPort = flag.Int("port", 0, "Server port, 0 means use 80 when disable HTTPS, 443 when enable")
	argPassword = flag.String("pwd", "rikka", "The password need provided when upload")
	argMaxSizeByMB = flag.Float64("size", 5, "Max file size by MB")
	argLogLevel = flag.Int(
		"level", logger.LevelInfo,
		fmt.Sprintf("Log level, from %d to %d", logger.LevelDebug, logger.LevelError),
	)
	argHTTPS = flag.Bool("https", false, "Use HTTPS")
	argCertDir = flag.String("certDir", ".", "Where to find HTTPS cert files(cert.pem, key.pem)")
	argAllowOrigin = flag.String("corsAllowOrigin", "", "Enable upload api CORS support, default is empty(disable). Set this to a origin, or * to enable for all origin")
	// Get name array of all available plugins, show in `rikka -h`
	pluginNames := make([]string, 0, len(pluginMap))
	for k := range pluginMap {
		pluginNames = append(pluginNames, k)
	}
	argPluginStr = flag.String(
		"plugin", "fs",
		"What plugin use to save file, selected from "+fmt.Sprintf("%v", pluginNames),
	)
}

func runtimeEnvCheck() {
	l.Info("Check runtime environment")

	l.Debug("Try to find plugin", *argPluginStr)

	// Make sure plugin be selected exist
	if plugin, ok := pluginMap[*argPluginStr]; ok {
		thePlugin = plugin
		l.Debug("Plugin", *argPluginStr, "found")
	} else {
		l.Fatal("Plugin", *argPluginStr, "not exist")
	}

	l.Info("All runtime environment check passed")
}
