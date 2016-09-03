package util

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/7sDream/rikka/common/logger"
)

var l = logger.NewLogger("[Util]")

// ErrHandle is a simple error handl function.
// If err is an error, write 500 InernalServerError to header and write error message to response and return true.
// Else (err is nil), don't do anything and return false.
func ErrHandle(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

// GetFilenameByRequest gets last part of url path as a filename and return it.
func GetTaskIDByRequest(r *http.Request) string {
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
		http.Error(w, "Method Not Allowed.", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

// Render is a shortcut function to render template to response.
func RenderTemplate(templatePath string, w http.ResponseWriter, data interface{}) error {
	t, err := template.ParseFiles(templatePath)
	if ErrHandle(w, err) {
		l.Warn("Parse template file", templatePath, "error:", err)
		return err
	}

	buff := bytes.NewBuffer([]byte{})

	err = t.Execute(buff, data)
	if ErrHandle(w, err) {
		l.Warn("Execute template", t, "with data", fmt.Sprintf("%+v", data), "error:", err)
		return err
	}

	content := make([]byte, buff.Len())
	buff.Read(content)
	w.Write(content)

	return nil
}

// RenderJSON is a shortcut function to write JSON data to response, and set the header Content-type
func RenderJSON(w http.ResponseWriter, data []byte) (err error) {
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	return err
}

// MustBeOr404 check if URL path is as excepted.
// If not equal, write 404 to header, "404 not fount" to response, and return false.
// Else don't do anything and return true.
func MustBeOr404(w http.ResponseWriter, r *http.Request, path string) bool {
	if r.URL.Path != path {
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
		l.Warn("Someone visit a non-exist page", r.URL.Path)
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
			l.Warn("Someone try to list dir", r.URL.Path)
			http.NotFound(w, r)
		} else {
			h.ServeHTTP(w, r)
		}
	}
}

type ContextCreator func(r *http.Request) interface{}

func TemplateRenderHandler(templatePath string, contextCreator ContextCreator, log *logger.Logger) http.HandlerFunc {
	if log == nil {
		log = l
	}
	return func(w http.ResponseWriter, r *http.Request) {
		defer recover()

		var err error
		if contextCreator != nil {
			err = RenderTemplate(templatePath, w, contextCreator(r))
		} else {
			err = RenderTemplate(templatePath, w, nil)
		}
		if err != nil {
			log.Warn("Render template", templatePath, "with data", nil, "error: ", err)
		}
	}
}

func RequestFilter(pathMustBe string, methodMustBe string, log *logger.Logger, handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer recover()

		if log == nil {
			log = l
		}

		if pathMustBe != "" {
			if !MustBeOr404(w, r, pathMustBe) {
				log.Warn("Someone visit a non-exist page", r.URL.Path, ", excepted is /")
				return
			}
		}

		if methodMustBe != "" {
			if !CheckMethod(w, r, methodMustBe) {
				log.Warn("Someone visit page", r.URL.Path, "with method", r.Method, ", only GET is allowed.")
				return
			}
		}

		handlerFunc(w, r)
	}
}
