package server

import (
	"net/http"

	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

var l = logger.NewLogger("[Rikka]")

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

func checkFiles() {
	l.Info("Check needed files")
	requireFiles := []string{
		"templates", "templates/index.html", "templates/view.html",
		"static", "static/main.css", "static/index.css", "static/view.css", "static/rikka.png",
	}

	for _, filepath := range requireFiles {
		if !util.CheckExist(filepath) {
			l.Fatal(filepath, "not exist, check failed, exit")
		}
	}
	l.Info("Check needed files successfully")
}

func StartRikkaServer(socket string, argPassword string, argMaxSizeByMB float64, plugin plugins.RikkaPlugin) {

	defer func() {
		recover()
		l.Info("Rikka stop")
	}()

	checkFiles()

	password = argPassword
	maxSizeByMB = argMaxSizeByMB

	staticFsHandler := http.StripPrefix(
		"/static",
		util.RequestFilter(
			"", "GET", l,
			util.DisableListDir(http.FileServer(http.Dir("static"))),
		),
	)

	indexHandler := util.RequestFilter(
		"/", "GET", l,
		util.TemplateRenderHandler("templates/index.html", nil, l),
	)

	viewHandler := util.RequestFilter(
		"", "GET", l,
		util.TemplateRenderHandler("templates/view.html", nil, l),
	)

	uploadHandler := util.RequestFilter(
		"/upload", "POST", l,
		upload,
	)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/view/", viewHandler)
	http.Handle("/static/", staticFsHandler)

	l.Info("Rikka main part started successfully, try to start plugin...")

	plugins.Load(plugin)

	l.Info("Rikka is listening", socket)

	err := http.ListenAndServe(socket, nil)
	if err != nil {
		l.Fatal(err.Error())
	}
}
