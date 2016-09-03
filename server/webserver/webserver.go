package webserver

import (
	"net/http"

	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

var l *logger.Logger

var port int
var password string
var maxSizeByMB float64

func upload(w http.ResponseWriter, r *http.Request) {
	defer recover()

	maxSize := int64(maxSizeByMB * 1024 * 1024)

	r.Body = http.MaxBytesReader(w, r.Body, maxSize)

	err := r.ParseMultipartForm(maxSize)
	if util.ErrHandle(w, err) {
		l.Error("Error happened when parse form:", err)
		return
	}

	userPassword := r.FormValue("password")
	if userPassword != password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error password."))
		l.Error("Someone input a error password:", userPassword)
		return
	}

	file, _, err := r.FormFile("uploadFile")
	if util.ErrHandle(w, err) {
		l.Error("Error happened when get form file:", err)
		return
	}

	TaskID, err := plugins.AcceptFile(&plugins.SaveRequest{File: file})
	if util.ErrHandle(w, err) {
		l.Error("Error happened when plugin revieve file save request:", err)
		return
	}

	w.Header().Set("Location", "/view/"+TaskID)
	w.WriteHeader(302)
}

// Check needed files like html, css, js, logo, etc...
func checkFiles() {
	requireFiles := []string{
		"templates", "templates/index.html", "templates/view.html",
		"static", "static/main.css", "static/index.css", "static/view.css", "static/rikka.png",
		"static/copy.js", "static/getSrc.js",
	}

	for _, filepath := range requireFiles {
		if !util.CheckExist(filepath) {
			l.Fatal(filepath, "not exist, check failed")
		}
	}

	l.Info("Check needed files successfully")
}

func viewHandleFunc(w http.ResponseWriter, r *http.Request) {
	taskID := util.GetTaskIDByRequest(r)

	pState, err := plugins.GetState(taskID)
	if util.ErrHandle(w, err) {
		return
	}

	if pState.StateCode == plugins.StateFinishCode {
		if url, err := plugins.GetURL(taskID, r, nil); err == nil {
			err = util.RenderTemplate("templates/viewFinish.html", w, url)
			util.ErrHandle(w, err)
		} else {
			util.ErrHandle(w, err)
		}
		return
	}

	err = util.RenderTemplate("templates/view.html", w, struct{ TaskID string }{TaskID: taskID})
	util.ErrHandle(w, err)
}

// StartRikkaWebServer start web server of rikka.
func StartRikkaWebServer(socket string, argPassword string, argMaxSizeByMB float64, log *logger.Logger) {

	l = log.SubLogger("[Web]")

	l.Info("Check needed files")
	checkFiles()

	// Save two important var to package level
	password = argPassword
	maxSizeByMB = argMaxSizeByMB

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

	// UploadHander handle request for upload photo(/upload)
	// Only accept POST Method
	uploadHandler := util.RequestFilter(
		"/upload", "POST", l,
		upload,
	)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/view/", viewHandler)
	http.Handle("/static/", staticFsHandler)

	l.Info("Rikka web server start successfully")
}
