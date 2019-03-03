package weibo

import (
	"flag"
	"net/http"

	"github.com/7sDream/rikka/plugins"
)

type weiboPlugin struct{}

var (
	l = plugins.SubLogger("[Weibo]")

	argUpdateCookiesPassword = flag.String(
		"ucpwd", "weibo",
		"Update cookies password, you need input this password when you visit /cookies to update your cookies",
	)

	client  *http.Client
	counter int64

	imageIDMap = make(map[int64]string)

	// Plugin is the main plugin object instance
	Plugin weiboPlugin
)

const (
	cookiesEnvKey = "RIKKA_WEIBO_COOKIES"
)
