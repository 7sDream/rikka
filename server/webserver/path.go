package webserver

// Can be configured
const (
	RootPath        = "/"
	ViewSuffix      = "view/"
	StaticSuffix    = "static/"
	faviconFileName = "favicon.png"
)

// Avaliable path of web server, DON'T configure thoose
const (
	ViewPath          = RootPath + ViewSuffix
	StaticPath        = RootPath + StaticSuffix
	FavIconOriginPath = "/favicon.ico"
	FavIconTruePath   = StaticPath + faviconFileName
)
