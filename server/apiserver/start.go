package apiserver

import (
	"net/http"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/common/util"
)

var (
	password    string
	maxSizeByMB float64

	l *logger.Logger
)

// StartRikkaAPIServer start API server of Rikka
func StartRikkaAPIServer(argPassword string, argMaxSizeByMb float64, log *logger.Logger) {

	password = argPassword
	maxSizeByMB = argMaxSizeByMb

	l = log.SubLogger("[API]")

	stateHandler := util.RequestFilter(
		"", "GET", l,
		util.DisableListDirFunc(l, stateHandleFunc),
	)

	urlHandler := util.RequestFilter(
		"", "GET", l,
		util.DisableListDirFunc(l, urlHandleFunc),
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
