package webserver

import "github.com/7sDream/rikka/api"

var (
	context = struct {
		Version     string
		RootPath    string
		UploadPath  string
		StaticPath  string
		FavIconPath string
		MaxSizeByMb float64
		TaskID      string
		URL         string
	}{
		Version:     api.Version,
		RootPath:    RootPath,
		UploadPath:  api.UploadPath,
		StaticPath:  StaticPath,
		FavIconPath: "",
		MaxSizeByMb: 0,
		TaskID:      "",
		URL:         "",
	}
)
