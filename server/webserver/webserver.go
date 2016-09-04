package webserver

import (
	"net/http"

	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

var l *logger.Logger

// Check needed files like html, css, js, logo, etc...
func checkFiles() {
	l.Info("Check needed files")

	requireFiles := []string{
		"templates", "templates/index.html", "templates/view.html", "templates/viewFinish.html",
		"static", "static/main.css", "static/index.css", "static/view.css", "static/rikka.png",
		"static/copy.js", "static/getSrc.js",
	}

	for _, filepath := range requireFiles {
		if !util.CheckExist(filepath) {
			l.Fatal(filepath, "not exist, check failed")
		} else {
			l.Debug("File", filepath, "exist, check passed")
		}
	}

	l.Info("Check needed files successfully")
}

func viewHandleFunc(w http.ResponseWriter, r *http.Request) {
	taskID := util.GetTaskIDByRequest(r)

	l.Debug("Recieve a view request of task", taskID)
	l.Debug("Send a state request of task", taskID, "to plugin")

	pState, err := plugins.GetState(taskID)
	if util.ErrHandle(w, err) {
		l.Warn("Get state of task", taskID, "error:", err)
		return
	}

	l.Debug("Get state of task", taskID, "successfully")

	if pState.StateCode == plugins.StateFinishCode {
		// state is finished
		templateFilePath := "templates/viewFinish.html"
		l.Debug("State of task", taskID, "is finished, render with", templateFilePath)
		l.Debug("Send an url request of task", taskID, "to plugin")
		if url, err := plugins.GetURL(taskID, r, nil); err == nil {
			// get url successfully
			l.Debug("Get url of task", taskID, "successfully:", url.URL)

			err = util.RenderTemplate(templateFilePath, w, url)

			if util.ErrHandle(w, err) {
				// RenderTemplate error
				l.Error("Error happened when render template", templateFilePath, ":", err)
			} else {
				// successfully
				l.Debug("Render template", templateFilePath, "successfully")
			}
			return
		}
		l.Error("Error happened when get url of task", taskID, ":", err)
		util.ErrHandle(w, err)
		return
	}

	// state is not finished
	templateFilePath := "templates/view.html"
	l.Warn("State of task", taskID, "is not finished(or error happened), render with", templateFilePath)

	err = util.RenderTemplate(templateFilePath, w, struct{ TaskID string }{TaskID: taskID})
	if util.ErrHandle(w, err) {
		// RenderTemplate error
		l.Error("Render template", templateFilePath, "error:", err)
	} else {
		// successfully
		l.Debug("Render template", templateFilePath, "successfully")
	}
}

// StartRikkaWebServer start web server of rikka.
func StartRikkaWebServer(log *logger.Logger) {

	l = log.SubLogger("[Web]")

	checkFiles()

	// The static file server handle all request that ask for files under /static
	// Only accept GET method
	staticFsHandler := http.StripPrefix(
		"/static",
		util.RequestFilter(
			"", "GET", l,
			util.DisableListDir(http.FileServer(http.Dir("static"))),
		),
	)

	// IndexHandler handle request ask for root(/, homepage of rikka), use templates/index.html
	// Only accept GET method.
	indexHandler := util.RequestFilter(
		"/", "GET", l,
		util.TemplateRenderHandler("templates/index.html", nil, l),
	)

	// ViewHandler handle requset ask for photo view page(/view/TaskID), use templates/view.html
	// Only accept GET Method
	viewHandler := util.RequestFilter(
		"", "GET", l,
		viewHandleFunc,
	)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/view/", viewHandler)
	http.Handle("/static/", staticFsHandler)

	l.Info("Rikka web server start successfully")
}
