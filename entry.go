package main

import (
	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/plugins"
	"github.com/7sDream/rikka/server"
)

// Logger of this package
var (
	l = logger.NewLogger("[Entry]")
)

// Main enterypoint
func main() {
	// print launch args
	l.Info(
		"Start rikka with arg:",
		"bind to socket", socket,
		", with password", *argPassword,
		", max file size", *argMaxSizeByMB, "MB",
		", plugin", *argPluginStr,
		", log level", *argLogLevel,
	)

	l.Info("Load plugin...")
	plugins.Load(thePlugin)

	// start Rikka servers (this call is Sync)
	server.StartRikka(socket, *argPassword, *argMaxSizeByMB)
}
