package weibo

import (
	"net/http"

	"github.com/7sDream/rikka/plugins"
)

type weiboPlugin struct{}

var (
	l = plugins.SubLogger("[Weibo]")

	client  *http.Client
	counter int64

	urlMap = make(map[int64]string)

	// WeiboPlugin is the main plugin object instance
	WeiboPlugin weiboPlugin
)

const (
	cookiesEnvKey = "RIKKA_WEIBO_COOKIES"
)

func (wbp weiboPlugin) ExtraHandlers() []plugins.HandlerWithPattern {
	return nil
}
