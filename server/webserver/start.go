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

	http.HandleFunc(RootPath, indexHandlerGenerator())
	http.HandleFunc(ViewPath, viewHandleGenerator())
	http.Handle(StaticPath, staticFsHandlerGenerator())

	l.Info("Rikka web server start successfully")
}
