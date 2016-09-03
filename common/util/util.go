package util

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/7sDream/rikka/common/logger"
)

var l = logger.NewLogger("[util]")

// ErrHandle is a simple error handl function.
// If err is an error, write 500 InernalServerError to header and write error message to response and return true.
// Else (err is nil), don't do anything and return false.
func ErrHandle(w http.ResponseWriter, err error) bool {
	if err != nil {
		l.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

// GetFilenameByRequest gets last part of url path as a filename and return it.
func GetFilenameByRequest(r *http.Request) string {
	splitedPath := strings.Split(r.URL.Path, "/")
	filename := splitedPath[len(splitedPath)-1]
	return filename
}

// CheckExist chekc a file or dir is Exist.
func CheckExist(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}
	return true
}

// CheckMethod check if request method is as excepted.
// If not, write the stauts "MethodNotAllow" to header, "Method Not Allowed." to response and return false.
// Else don't do anything and return true.
func CheckMethod(w http.ResponseWriter, r *http.Request, excepted string) bool {
	if r.Method != excepted {
		l.Error("Someone visit page", r.URL.Path, "with method", r.Method, ", only", excepted, "is allowed.")
		http.Error(w, "Method Not Allowed.", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

// Render is a shortcut function to render template to response.
func Render(templatePath string, w http.ResponseWriter, data interface{}) {
	t, err := template.ParseFiles(templatePath)
	if ErrHandle(w, err) {
		return
	}

	buff := bytes.NewBuffer([]byte{})

	err = t.Execute(buff, data)
	if ErrHandle(w, err) {
		return
	}
	content := make([]byte, buff.Len())
	buff.Read(content)
	w.Write(content)
}

// MustBeOr404 check if URL path is as excepted.
// If not equal, write 404 to header, "404 not fount" to response, and return false.
// Else don't do anything and return true.
func MustBeOr404(w http.ResponseWriter, r *http.Request, path string) bool {
	if r.URL.Path != path {
		l.Error("Someone visit a non-exist page", r.URL.Path, ", excepted is", path)
		http.NotFound(w, r)
		return false
	}
	return true
}

// MustExistOr404 check if a file is exist.
// If not, write 404 to header, "404 not fount" to response, and return false.
// Else don't do anything and return true.
func MustExistOr404(w http.ResponseWriter, r *http.Request, filepath string) bool {
	if !CheckExist(filepath) {
		l.Error("Someone visit a non-exist page", r.URL.Path)
		http.NotFound(w, r)
		return false
	}
	return true
}

// DisableListDir accept a FileServer handle and return a handle that not allow
// list dir.
func DisableListDir(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			l.Error("Someone try to list dir", r.URL.Path)
			http.NotFound(w, r)
		} else {
			h.ServeHTTP(w, r)
		}
	}
}
