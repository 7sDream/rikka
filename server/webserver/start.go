package webserver

import (
	"net/http"

	"github.com/7sDream/rikka/common/logger"
)

var l *logger.Logger

// StartRikkaWebServer start web server of rikka.
func StartRikkaWebServer(log *logger.Logger) {

	l = log.SubLogger("[Web]")

	checkFiles()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/view/", viewHandler)
	http.Handle("/static/", staticFsHandler)

	l.Info("Rikka web server start successfully")
}
