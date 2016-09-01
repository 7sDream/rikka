package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func errHandle(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("templates/index.html")
		if err == nil {
			err = t.Execute(w, nil)
			if err == nil {
				return
			}
		}
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte("Error Method."))
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	defer func() {
		recover()
	}()

	if r.Method == "POST" {
		err := r.ParseMultipartForm(1024 * 5)
		errHandle(w, err)
		file, handle, err := r.FormFile("uploadFile")
		errHandle(w, err)
		defer file.Close()
		saveFile, err := ioutil.TempFile("files", handle.Filename)
		errHandle(w, err)
		defer saveFile.Close()
		io.Copy(saveFile, file)
		w.Header().Set("Location", "/"+saveFile.Name())
		w.WriteHeader(302)
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error Method."))
	}
}

func files(w http.ResponseWriter, r *http.Request) {
	defer func() {
		recover()
	}()

	if r.Method == "GET" {
		filepath := r.URL.Path[1:]
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		content, err := ioutil.ReadFile(filepath)
		errHandle(w, err)
		w.Write(content)
	}
}

func main() {
	if _, err := os.Stat("files"); os.IsNotExist(err) {
		os.MkdirAll("files", os.ModeDir)
	}
	http.HandleFunc("/index", index)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/files/", files)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
