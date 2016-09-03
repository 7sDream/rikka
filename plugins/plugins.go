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
	return currentPlugin.SaveRequestHandle(q)
}

func GetState(taskID string) (r *State, err error) {
	return currentPlugin.StateRequestHandle(taskID)
}

func GetURL(taskID string, r *http.Request, picOp *PictureOperate) (url *URL, err error) {
	return currentPlugin.GetSrcURL(&SrcURLRequest{
		HttpRequest: r,
		TaskID:      taskID,
		PicOp:       picOp,
	})
}
