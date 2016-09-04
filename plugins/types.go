package plugins

import (
	"mime/multipart"
	"net/http"
)

// SaveRequest is a request that want to 'save'(acctuly upload) a file.
// Plugins' SaveRequestHandle func should accept a point of instance and return a string as taskID
type SaveRequest struct {
	File multipart.File
}

// URLRequest is a request ask for photo src url of a task.
// Plugins' URLRequestHandle func should accept a point of instance and return a string as URL
type URLRequest struct {
	HTTPRequest *http.Request
	TaskID      string
	PicOp       *ImageOperate
}

// State shows a state of task.
type State struct {
	TaskID      string
	StateCode   int
	State       string
	Description string
}

// ImageOperate stand for some operate of src imgage, not used now.
type ImageOperate struct {
	Width    int
	Height   int
	Rotate   int
	OtherArg string
}

// ErrorJSON struct used to build json from eror string.
type ErrorJSON struct {
	Error string
}

// URLJSON struct used to build json from eror URL.
type URLJSON struct {
	URL string
}

// TaskIDJSON struct used to build json from taskID.
type TaskIDJSON struct {
	TaskID string
}

// HandlerWithPattern is a struct combine a http.Handler with the pattern is will server.
// Plugins' ExtraHandlers func return an array of this and will be added to http.handler when init plugin.
type HandlerWithPattern struct {
	Pattern string
	Handler http.Handler
}

// RikkaPlugin is plugin interface, all plugin should implement thoose function.
type RikkaPlugin interface {
	// Init will be called when load plugin.
	Init()
	// AcceptFile will call this.
	SaveRequestHandle(*SaveRequest) (string, error)
	// GetState will call this.
	StateRequestHandle(string) (*State, error)
	// GetURL will call this.
	URLRequestHandle(q *URLRequest) (*URLJSON, error)
	// Will be added into http handler list.
	ExtraHandlers() []HandlerWithPattern
}
