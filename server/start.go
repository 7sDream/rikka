package server

import (
	"net/http"

	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/plugins"
	"github.com/7sDream/rikka/server/apiserver"
	"github.com/7sDream/rikka/server/webserver"
)

var l = logger.NewLogger("[Server]")

// StartRikka start all part of rikka. Include File process plugin, web Server and API server.
func StartRikka(socket string, password string, maxSizeByMB float64, plugin plugins.RikkaPlugin) {

	l.Info("Load plugin...")
	plugins.Load(plugin)

	l.Info("Start API server")
	apiserver.StartRikkaAPIServer(password, maxSizeByMB, l)

	l.Info("Start web server...")
	webserver.StartRikkaWebServer(l)

	l.Info("Rikka is listening", socket)

	// real http server function call
	err := http.ListenAndServe(socket, nil)

	if err != nil {
		l.Fatal("Error when try listening", socket, ":", err)
	}
}
