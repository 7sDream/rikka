package webserver

// Can be configured
const (
	RootPath        = "/"
	ViewSuffix      = "view/"
	StaticSuffix    = "static/"
	faviconFileName = "image/favicon.png"
)

// Available path of web server, DON'T configure those
const (
	ViewPath          = RootPath + ViewSuffix
	StaticPath        = RootPath + StaticSuffix
	FavIconOriginPath = "/favicon.ico"
	FavIconTruePath   = StaticPath + faviconFileName
)
