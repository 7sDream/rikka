package webserver

import (
	"net/http"

	"github.com/7sDream/rikka/common/util"
)

// IndexHandler handle request ask for root(/, homepage of rikka), use templates/index.html
// Only accept GET method.
var indexHandler = util.RequestFilter(
	"/", "GET", l,
	util.TemplateRenderHandler(
		"templates/index.html",
		func(r *http.Request) interface{} { return context }, l,
	),
)
