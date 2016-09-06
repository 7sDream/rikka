package webserver

import (
	"net/http"

	"github.com/7sDream/rikka/common/util"
)

// IndexHandler handle request ask for homepage(${RootPath}, "/" in general), use templates/index.html
// Only accept GET method.
func indexHandlerGenerator() http.HandlerFunc {
	return util.RequestFilter(
		RootPath, "GET", l,
		util.TemplateRenderHandler(
			homeTemplateFilePath,
			func(r *http.Request) interface{} { return context }, l,
		),
	)
}
