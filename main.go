package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	pathutil "path/filepath"
	"strconv"
	"time"

	"github.com/7sDream/rikka/util"
)

var port = flag.Int("port", 80, "server port")
var password = flag.String("pwd", "rikka", "the password")

type viewPhoto struct {
	Filename string
	URL      string
}

func buildURL(r *http.Request) string {
	res := url.URL{
		Scheme: "http",
		Host:   r.Host,
		Path:   "files/" + util.GetFilenameByRequest(r),
	}
	return res.String()
}

func index(w http.ResponseWriter, r *http.Request) {
	defer recover()

	if !util.MustBeOr404(w, r, "/") {
		return
	}

	if !util.CheckMethod(w, r, "GET") {
		return
	}

	util.Render("templates/index.html", w, nil)
}

func upload(w http.ResponseWriter, r *http.Request) {
	defer recover()

	if !util.MustBeOr404(w, r, "/upload") {
		return
	}

	if !util.CheckMethod(w, r, "POST") {
		return
	}

	err := r.ParseMultipartForm(1024 * 5)
	if util.ErrHandle(w, err) {
		return
	}

	userPassword := r.FormValue("password")
	if userPassword != *password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error password."))
		util.Error("Someont input a error password:", userPassword)
		return
	}

	file, _, err := r.FormFile("uploadFile")
	if util.ErrHandle(w, err) {
		return
	}
	defer file.Close()

	now := time.Now().UTC()
	saveFile, err := ioutil.TempFile("files", now.Format("2006-01-02-"))
	if util.ErrHandle(w, err) {
		return
	}
	defer saveFile.Close()

	io.Copy(saveFile, file)

	_, name := pathutil.Split(saveFile.Name())
	w.Header().Set("Location", "/view/"+name)
	w.WriteHeader(302)

	util.Info("Accepted a new file:", name)
}

func view(w http.ResponseWriter, r *http.Request) {
	defer recover()

	if !util.CheckMethod(w, r, "GET") {
		return
	}

	filename := util.GetFilenameByRequest(r)
	filepath := "files/" + filename
	if !util.MustExistOr404(w, r, filepath) {
		util.Error("Someone visit a non-exist photo:", filename)
		return
	}

	view := viewPhoto{
		Filename: filename,
		URL:      buildURL(r),
	}

	util.Render("templates/view.html", w, view)
}

func main() {

	flag.Parse()

	util.Info("Args port =", *port)
	util.Info("Args password =", *password)

	requireFiles := []string{
		"files",
		"templates", "templates/index.html", "templates/view.html",
		"static", "static/main.css", "static/index.css", "static/view.css", "static/rikka.png",
	}

	for _, file := range requireFiles {
		if !util.CheckExist(file) {
			util.Error(file, "not exist, exit.")
			return
		}
	}

	staticFs := util.DisableListDir(http.FileServer(http.Dir("static")))
	fileFs := util.DisableListDir(http.FileServer(http.Dir("files")))

	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/view/", view)
	http.Handle("/files/", http.StripPrefix("/files", fileFs))
	http.Handle("/static/", http.StripPrefix("/static", staticFs))

	util.Info("Rikka started")

	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
