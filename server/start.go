package server

import (
	"net/http"
	pathUtil "path/filepath"

	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/server/apiserver"
	"github.com/7sDream/rikka/server/webserver"
)

var (
	l = logger.NewLogger("[Server]")
)

// StartRikka start server part of rikka. Include Web Server and API server.
func StartRikka(socket string, password string, maxSizeByMb float64, https bool, certDir string) {
	realHttps := false

	certPemPath := pathUtil.Join(certDir, "cert.pem")
	keyPemPath := pathUtil.Join(certDir, "key.pem")

	if https {
		if util.IsFile(certPemPath) && util.IsFile(keyPemPath) {
			realHttps = true
		} else {
			l.Warn("Cert dir argument is not a valid dir, fallback to http")
		}
	}

	l.Info("Start web server...")
	viewPath := webserver.StartRikkaWebServer(maxSizeByMb, https, l)

	l.Info("Start API server...")
	apiserver.StartRikkaAPIServer(viewPath, password, maxSizeByMb, https, l)

	l.Info("Rikka is listening", socket)

	// real http server function call
	var err error
	if realHttps {
		err = http.ListenAndServeTLS(socket, certPemPath, keyPemPath, nil)
	} else {
		err = http.ListenAndServe(socket, nil)
	}

	if err != nil {
		l.Fatal("Error when try listening", socket, ":", err)
	}
}
