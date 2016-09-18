package weibo

import (
	"errors"
	"net/http"

	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

const (
	updateCookiesFormHTML = `
    <!doctype html>
    <html>
        <head>
            <title>Update cookies</title>
        </head>
        <body>
            <form id="updateCookies" method="POST" action="/update">
                <textarea cols="80" rows="10" form="updateCookies" name="cookies"></textarea><br>
                <input type="password" name="password"><br>
                <input type="submit" >
            </form>
        </body>
    </html>
    `
)

func (wbp weiboPlugin) ExtraHandlers() []plugins.HandlerWithPattern {
	updateCookiesFormHandler := plugins.HandlerWithPattern{
		Pattern: "/cookies",
		Handler: util.RequestFilter(
			"/cookies", "GET", l,
			util.TemplateStringRenderHandler(
				"cookiesForm.html", updateCookiesFormHTML, nil, l,
			),
		),
	}

	updateCookiesHander := plugins.HandlerWithPattern{
		Pattern: "/update",
		Handler: util.RequestFilter(
			"/update", "POST", l,
			func(w http.ResponseWriter, r *http.Request) {
				if r.FormValue("password") != *argUpdateCookiesPassword {
					util.ErrHandle(w, errors.New("Error password"))
					return
				}
				err := updateCookies(r.FormValue("cookies"))
				if util.ErrHandle(w, err) {
					return
				}
				http.Redirect(w, r, "/", http.StatusFound)
			},
		),
	}

	return []plugins.HandlerWithPattern{updateCookiesFormHandler, updateCookiesHander}
}
