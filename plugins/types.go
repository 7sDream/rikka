package plugins

import (
	"mime/multipart"
	"net/http"
)

// SaveRequest is a file save request.
type SaveRequest struct {
	File multipart.File
}

type SrcURLRequest struct {
	HTTPRequest *http.Request
	TaskID      string
	PicOp       *PictureOperate
}

type State struct {
	TaskID      string
	StateCode   int
	State       string
	Description string
}

type PictureOperate struct {
	Width    int
	Height   int
	Rotate   int
	OtherArg string
}

type Error struct {
	Error string
}

type URL struct {
	URL string
}

type TaskID struct {
	TaskID string
}

type HandlerWithPattern struct {
	Pattern string
	Handler http.Handler
}

type RikkaPlugin interface {
	Init()
	SaveRequestHandle(*SaveRequest) (string, error)
	StateRequestHandle(string) (*State, error)
	GetSrcURL(q *SrcURLRequest) (*URL, error)
	ExtraHandlers() []HandlerWithPattern
}
