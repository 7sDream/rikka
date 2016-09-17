package plugins

import (
	"mime/multipart"
	"net/http"

	"github.com/7sDream/rikka/api"
)

// SaveRequest is a request that want to 'save'(acctuly upload) a file.
// Plugins' SaveRequestHandle func should accept a point of instance and return a string as taskID
type SaveRequest struct {
	File     multipart.File
	FileSize int64
	FileExt  string
}

// URLRequest is a request ask for photo src url of a task.
// Plugins' URLRequestHandle func should accept a point of instance and return a string as URL
type URLRequest struct {
	HTTPRequest *http.Request
	TaskID      string
	PicOp       *ImageOperate
}

// ImageOperate stand for some operate of src imgage, not used now.
type ImageOperate struct {
	Width    int
	Height   int
	Rotate   int
	OtherArg string
}

// HandlerWithPattern is a struct combine a http.Handler with the pattern is will server.
// Plugins' ExtraHandlers func return an array of this and will be added to http.handler when init plugin.
type HandlerWithPattern struct {
	Pattern string
	Handler http.HandlerFunc
}

// RikkaPlugin is plugin interface, all plugin should implement thoose function.
type RikkaPlugin interface {
	// Init will be called when load plugin.
	Init()
	// AcceptFile will call this.
	SaveRequestHandle(*SaveRequest) (*api.TaskID, error)
	// GetState will call this.
	StateRequestHandle(string) (*api.State, error)
	// GetURL will call this.
	URLRequestHandle(q *URLRequest) (*api.URL, error)
	// Will be added into http handler list.
	ExtraHandlers() []HandlerWithPattern
}
