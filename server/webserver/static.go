package webserver

import (
	"net/http"

	"github.com/7sDream/rikka/common/util"
)

// The static file server handle all request that ask for files under /static
// Only accept GET method
var staticFsHandler = http.StripPrefix(
	"/static",
	util.RequestFilter(
		"", "GET", l,
		util.DisableListDir(http.FileServer(http.Dir("static"))),
	),
)
