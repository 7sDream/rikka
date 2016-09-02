package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	pathutil "path/filepath"
	"strconv"
	"time"

	"github.com/7sDream/rikka/util"
)

var password string
var defaultDomain = "localhost"

const port int = 8000

type viewPhoto struct {
	Filename string
	URL      string
}

func buildURL(r *http.Request) string {
	host := r.Host
	if host == "" {
		host = defaultDomain
		if port != 80 {
			host += ":80"
		}
	}
	res := url.URL{
		Scheme: "http",
		Host:   host,
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
	if userPassword != password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error password."))
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
}

func view(w http.ResponseWriter, r *http.Request) {
	defer recover()

	if !util.CheckMethod(w, r, "GET") {
		return
	}

	filename := util.GetFilenameByRequest(r)
	filepath := "files/" + filename
	if !util.CheckExist(filepath) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No this photo."))
	}

	view := viewPhoto{
		Filename: filename,
		URL:      buildURL(r),
	}

	util.Render("templates/view.html", w, view)
}

func main() {
	requireFiles := []string{
		"files",
		"templates", "templates/index.html", "templates/view.html",
		"static", "static/main.css", "static/index.css", "static/view.css", "static/rikka.png",
	}

	for _, file := range requireFiles {
		if !util.CheckExist(file) {
			fmt.Println(file, "not exist, exit.")
			return
		}
	}

	password = os.Getenv("RIKKA_PWD")

	if password == "" {
		fmt.Println("No password proivede, use [rikka] as password.")
		password = "rikka"
	} else {
		fmt.Println("Get password from env: ", password)
	}

	staticFs := http.FileServer(http.Dir("static"))
	fileFs := http.FileServer(http.Dir("files"))

	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/view/", view)
	http.Handle("/files/", http.StripPrefix("/files/", fileFs))
	http.Handle("/static/", http.StripPrefix("/static/", staticFs))

	fmt.Println("Starting rikka...")

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
