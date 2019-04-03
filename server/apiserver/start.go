package apiserver

import (
	"net/http"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/common/util"
)

var (
	password    string
	maxSizeByMb float64
	isServerTLS bool

	l *logger.Logger
)

// StartRikkaAPIServer start API server of Rikka
func StartRikkaAPIServer(argViewPath string, argPassword string, argMaxSizeByMb float64, argIsServerTLS bool, log *logger.Logger) {

	viewPath = argViewPath
	password = argPassword
	maxSizeByMb = argMaxSizeByMb
	isServerTLS = argIsServerTLS

	l = log.SubLogger("[API]")

	stateHandler := util.RequestFilter(
		"", "GET", l,
		util.DisableListDir(l, stateHandleFunc),
	)

	urlHandler := util.RequestFilter(
		"", "GET", l,
		util.DisableListDir(l, urlHandleFunc),
	)

	uploadHandler := util.RequestFilter(
		api.UploadPath, "POST", l,
		uploadHandleFunc,
	)

	http.HandleFunc(api.StatePath, stateHandler)
	http.HandleFunc(api.URLPath, urlHandler)
	http.HandleFunc(api.UploadPath, uploadHandler)

	l.Info("API server start successfully")

	return
}
