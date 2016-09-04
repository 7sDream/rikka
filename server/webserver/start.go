package webserver

import (
	"net/http"

	"github.com/7sDream/rikka/common/logger"
)

var (
	l *logger.Logger
)

// StartRikkaWebServer start web server of rikka.
func StartRikkaWebServer(log *logger.Logger) {

	l = log.SubLogger("[Web]")

	checkFiles()

	http.HandleFunc(RootPath, indexHandler)
	http.HandleFunc(ViewPath, viewHandler)
	http.Handle(StaticPath, staticFsHandler)

	l.Info("Rikka web server start successfully")
}
