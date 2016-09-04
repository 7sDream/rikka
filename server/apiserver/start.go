package apiserver

import (
	"net/http"

	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/common/util"
)

var password string
var maxSizeByMB float64

var l *logger.Logger

// StartRikkaAPIServer start API server of Rikka
func StartRikkaAPIServer(argPassword string, argMaxSizeByMb float64, log *logger.Logger) {

	password = argPassword
	maxSizeByMB = argMaxSizeByMb

	l = log.SubLogger("[API]")

	stateHandler := util.RequestFilter(
		"", "GET", l,
		util.DisableListDirFunc(stateHandleFunc),
	)

	urlHandler := util.RequestFilter(
		"", "GET", l,
		util.DisableListDirFunc(urlHandleFunc),
	)

	uploadHandler := util.RequestFilter(
		"/api/upload", "POST", l,
		uploadHandleFunc,
	)

	http.HandleFunc("/api/state/", stateHandler)
	http.HandleFunc("/api/url/", urlHandler)
	http.HandleFunc("/api/upload", uploadHandler)

	l.Info("API server start successfully")

	return
}
