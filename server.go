package main

import (
	"net/http"

	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

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

	maxSize := int64(*argMaxSizeByMB * 1024 * 1024)

	r.Body = http.MaxBytesReader(w, r.Body, maxSize)

	err := r.ParseMultipartForm(maxSize)
	if util.ErrHandle(w, err) {
		return
	}

	userPassword := r.FormValue("password")
	if userPassword != *argPassword {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error password."))
		l.Error("Someone input a error password:", userPassword)
		return
	}

	file, _, err := r.FormFile("uploadFile")
	if util.ErrHandle(w, err) {
		return
	}

	fileID, err := plugins.AcceptFile(&plugins.SaveRequest{File: file})
	if util.ErrHandle(w, err) {
		return
	}

	w.Header().Set("Location", "/view/"+fileID)
	w.WriteHeader(302)
}

func view(w http.ResponseWriter, r *http.Request) {
	defer recover()

	if !util.CheckMethod(w, r, "GET") {
		return
	}

	util.Render("templates/view.html", w, nil)
}
