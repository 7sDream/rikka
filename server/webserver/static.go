package webserver

import (
	"net/http"

	"github.com/7sDream/rikka/common/util"
)

// The static file server handle all request that ask for files under static dir, from url path {StaticPath}<filename>
// Only accept GET method
func staticFsHandlerGenerator() http.Handler {
	return util.RequestFilter(
		"", "GET", l,
		util.DisableListDir(
			l,
			http.StripPrefix(
				StaticPath[:len(StaticPath)-1],
				http.FileServer(http.Dir("static")),
			),
		),
	)
}
