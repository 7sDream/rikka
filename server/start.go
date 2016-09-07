package server

import (
	"net/http"

	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/server/apiserver"
	"github.com/7sDream/rikka/server/webserver"
)

var (
	l = logger.NewLogger("[Server]")
)

// StartRikka start server part of rikka. Include Web Server and API server.
func StartRikka(socket string, password string, maxSizeByMb float64) {

	l.Info("Start web server...")
	viewPath := webserver.StartRikkaWebServer(maxSizeByMb, l)

	l.Info("Start API server")
	apiserver.StartRikkaAPIServer(viewPath, password, maxSizeByMb, l)

	l.Info("Rikka is listening", socket)

	// real http server function call
	err := http.ListenAndServe(socket, nil)

	if err != nil {
		l.Fatal("Error when try listening", socket, ":", err)
	}
}
