package webserver

import (
	"net/http"

	"github.com/7sDream/rikka/common/logger"
)

var (
	l *logger.Logger
)

// StartRikkaWebServer start web server of rikka.
func StartRikkaWebServer(maxSizeByMb float64, log *logger.Logger) {

	if maxSizeByMb <= 0 {
		l.Fatal("Max file size can't be equal or less than 0, you set it to", maxSizeByMb)
	}

	context.MaxSizeByMb = maxSizeByMb

	l = log.SubLogger("[Web]")

	checkFiles()

	http.HandleFunc(RootPath, indexHandlerGenerator())
	http.HandleFunc(ViewPath, viewHandleGenerator())
	http.HandleFunc(StaticPath, staticFsHandlerGenerator())
	http.HandleFunc(FavIconOriginPath, favIconHandlerGenerator())

	l.Info("Rikka web server start successfully")
}
