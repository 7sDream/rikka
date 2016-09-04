package fs

import (
	"net/http"

	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

// ExtraHandlers return value will be add to http handle list.
// In fs plugin, we start a static file server to serve image file we accped in /files/taskID path.
func (fsp fsPlugin) ExtraHandlers() (handlers []plugins.HandlerWithPattern) {
	// get a base file server
	fileServer := http.StripPrefix(
		"/files",
		// Disable list dir
		util.DisableListDir(
			http.FileServer(http.Dir(imageDir)),
		),
	)
	// only accped GET request
	requestFilterFileServer := util.RequestFilter(
		"", "GET", l,
		func(w http.ResponseWriter, q *http.Request) {
			fileServer.ServeHTTP(w, q)
		},
	)

	handlers = []plugins.HandlerWithPattern{
		plugins.HandlerWithPattern{
			Pattern: "/files/", Handler: requestFilterFileServer,
		},
	}

	return handlers
}
