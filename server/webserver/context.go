package webserver

import "github.com/7sDream/rikka/api"

var (
	context = struct {
		Version    string
		RootPath   string
		UploadPath string
		StaticPath string
		TaskID     string
		URL        string
	}{
		Version:    api.Version,
		RootPath:   RootPath,
		UploadPath: api.UploadPath,
		StaticPath: StaticPath,
		TaskID:     "",
		URL:        "",
	}
)
