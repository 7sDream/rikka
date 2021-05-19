package apiserver

import (
	"net/http"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/common/util"
)

var (
	password        string
	maxSizeByMb     float64
	isServerTLS     bool
	corsAllowOrigin string

	l *logger.Logger
)

func addCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(corsAllowOrigin) > 0 {
			w.Header().Set("Access-Control-Allow-Origin", corsAllowOrigin)
		}
		h(w, r)
	}
}

// StartRikkaAPIServer start API server of Rikka
func StartRikkaAPIServer(argViewPath string, argPassword string, argMaxSizeByMb float64, argIsServerTLS bool, argCorsAllowOrigin string, log *logger.Logger) {

	viewPath = argViewPath
	password = argPassword
	maxSizeByMb = argMaxSizeByMb
	isServerTLS = argIsServerTLS
	corsAllowOrigin = argCorsAllowOrigin

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

	http.HandleFunc(api.StatePath, addCors(stateHandler))
	http.HandleFunc(api.URLPath, addCors(urlHandler))
	http.HandleFunc(api.UploadPath, addCors(uploadHandler))

	l.Info("API server start successfully")
}
