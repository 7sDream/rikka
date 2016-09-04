package plugins

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/common/logger"
)

var (
	l = logger.NewLogger("[Plugins]")

	currentPlugin RikkaPlugin
)

// SubLogger return a new sub logger from plugins logger.
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
func AcceptFile(q *SaveRequest) (fileID *api.TaskID, err error) {
	return currentPlugin.SaveRequestHandle(q)
}

// GetState will be called when API server recieve a state request.
// Also be called when web server recieve a view request,
// web server decide response a finished view html or a self-renewal html based on
// the return state is finished state.
func GetState(taskID string) (r *api.State, err error) {
	return currentPlugin.StateRequestHandle(taskID)
}

// GetURL will be called when API server recieve a url request.
// Also be called when web server recieve a view request and GetState return a finished state.
// web server use the return url value to render a finished view html.
func GetURL(taskID string, r *http.Request, picOp *ImageOperate) (pURL *api.URL, err error) {
	l.Debug("Send state request to plugin before get url of task", taskID)
	var pState *api.State

	// check state successfully
	if pState, err = GetState(taskID); err == nil {
		// not finished
		if pState.StateCode != api.StateFinishCode {
			l.Warn("Task", taskID, "not finished, can't get url")
			return nil, errors.New("Task not finished")
		}
		// finished
		l.Debug("Task", taskID, "is finished, send url request to the plugin")
		return currentPlugin.URLRequestHandle(&URLRequest{
			HTTPRequest: r,
			TaskID:      taskID,
			PicOp:       picOp,
		})
	}

	// check state error
	errorMsg := fmt.Sprint("Error happened when get state of task", taskID, ":", err)
	l.Error(errorMsg)
	return nil, errors.New(errorMsg)
}
