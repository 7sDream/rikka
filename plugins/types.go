package plugins

import (
	"mime/multipart"
	"net/http"
)

// SaveRequest is a file save request.
type SaveRequest struct {
	File multipart.File
}

// SaveResponse isa response of SaveRequest.
type SaveResponse struct {
	TaskID string
}

type State struct {
	TaskID      string
	StateCode   int
	State       string
	Description string
}

// StateRequest is a request to get state of a task.
type StateRequest struct {
	TaskID string
}

// StateResponse is response of StateRequest.
type StateResponse struct {
	State State
}

type PictureOperate struct {
	Width    int
	Height   int
	Rotate   int
	OtherArg string
}

type SrcURLRequest struct {
	HttpRequest *http.Request
	TaskID      string
	Operate     PictureOperate
}

type HandlerWithPattern struct {
	Pattern string
	Handler http.Handler
}

type RikkaPlugin interface {
	Init()
	SaveRequestHandle(q *SaveRequest) (response *SaveResponse, err error)
	StateRequestHandle(q *StateRequest) (response *StateResponse, err error)
	GetSrcURL(q *SrcURLRequest) (url string, err error)
	ExtraHandlers() []HandlerWithPattern
}
