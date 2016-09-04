package fs

import (
	"net/http"

	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

const (
	fileURLPath = "/files/"
)

// ExtraHandlers return value will be add to http handle list.
// In fs plugin, we start a static file server to serve image file we accped in /files/taskID path.
func (fsp fsPlugin) ExtraHandlers() (handlers []plugins.HandlerWithPattern) {
	// get a base file server
	fileServer := http.StripPrefix(
		fileURLPath[:len(fileURLPath)-1],
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
			Pattern: fileURLPath, Handler: requestFilterFileServer,
		},
	}

	return handlers
}
