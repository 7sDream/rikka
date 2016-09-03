package plugins

import (
	"net/http"
	"sync"

	"github.com/7sDream/rikka/common/logger"
)

var l = logger.NewLogger("[Plugins]")
var tasks = struct {
	sync.RWMutex
	m map[string]*State
}{m: make(map[string]*State)}

var currentPlugin RikkaPlugin

// SubLogger return a new sub logger belongs to plugins logger.
func SubLogger(prefix string) *logger.Logger {
	return l.SubLogger(prefix)
}

// Load load a plugin to net/http
func Load(plugin RikkaPlugin) {
	currentPlugin = plugin
	currentPlugin.Init()
	for _, hp := range currentPlugin.ExtraHandlers() {
		http.Handle(hp.Pattern, hp.Handler)
	}
}

// AcceptFile whill be called when you recieve a file upload request, the SaveRequest struct contains the file.
func AcceptFile(q *SaveRequest) (fileID string, err error) {
	res, err := currentPlugin.SaveRequestHandle(q)
	if err == nil {
		return res.TaskID, nil
	}
	return "", err
}
