package webserver

import (
	"net/http"

	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/common/util"
)

var (
	isServeTLS bool
	l *logger.Logger
)

// StartRikkaWebServer start web server of rikka.
func StartRikkaWebServer(maxSizeByMb float64, argIsServeTLS bool, log *logger.Logger) string {

	if maxSizeByMb <= 0 {
		l.Fatal("Max file size can't be equal or less than 0, you set it to", maxSizeByMb)
	}

	isServeTLS = argIsServeTLS

	// change all sub-folder in content
	subFolder := util.GetSubFolder()
	if len(subFolder) > 1 {
		tmpStr := subFolder[:len(subFolder) - 1]
		context.RootPath = tmpStr + "/"
		context.StaticPath = tmpStr + context.StaticPath
		context.UploadPath = tmpStr + context.UploadPath
	}

	context.MaxSizeByMb = maxSizeByMb
	context.FavIconPath = FavIconTruePath

	l = log.SubLogger("[Web]")

	checkFiles()

	http.HandleFunc(RootPath, indexHandlerGenerator())
	http.HandleFunc(ViewPath, viewHandleGenerator())
	http.HandleFunc(StaticPath, staticFsHandlerGenerator())
	http.HandleFunc(FavIconOriginPath, favIconHandlerGenerator())

	l.Info("Rikka web server start successfully")

	return ViewPath
}
