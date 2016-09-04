package webserver

import "github.com/7sDream/rikka/api"

var (
	context = struct {
		Version    string
		TaskID     string
		URL        string
		RootPath   string
		UploadPath string
		StaticPath string
	}{
		Version:    "0.1.1",
		TaskID:     "",
		URL:        "",
		RootPath:   RootPath,
		UploadPath: api.UploadPath,
		StaticPath: StaticPath,
	}
)
