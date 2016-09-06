package util

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/7sDream/rikka/common/logger"
)

var (
	l = logger.NewLogger("[Util]")
)

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

// GetTaskIDByRequest gets last part of url path as a taskID and return it.
func GetTaskIDByRequest(r *http.Request) string {
	splitedPath := strings.Split(r.URL.Path, "/")
	filename := splitedPath[len(splitedPath)-1]
	return filename
}

// CheckExist chekc if a file or dir is Exist.
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
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

// RenderTemplate is a shortcut function to render template to response.
func RenderTemplate(templatePath string, w http.ResponseWriter, data interface{}) error {
	t, err := template.ParseFiles(templatePath)
	if ErrHandle(w, err) {
		l.Error("Error happened when parse template file", templatePath, ":", err)
		return err
	}

	// a buff that ues in execute, if error happened,
	// error message will not be write to truly response
	buff := bytes.NewBuffer([]byte{})
	err = t.Execute(buff, data)

	// error happened, write a generic error message to response
	if err != nil {
		l.Error("Error happened when execute template", t, "with data", fmt.Sprintf("%+v", data), ":", err)
		ErrHandle(w, errors.New("error when render template"))
		return err
	}

	// no error happened, write to response
	content := make([]byte, buff.Len())
	buff.Read(content)
	_, err = w.Write(content)

	return err
}

// RenderJSON is a shortcut function to write JSON data to response, and set the header Content-Type.
func RenderJSON(w http.ResponseWriter, data []byte, code int) (err error) {
	w.WriteHeader(code)
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
		http.NotFound(w, r)
		return false
	}
	return true
}

// DisableListDir accept a FileServer handle and return a handle that not allow list dir.
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

// DisableListDirFunc accept a handle func and return a handle that not allow list dir.
func DisableListDirFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			l.Warn("Someone try to list dir", r.URL.Path)
			http.NotFound(w, r)
		} else {
			h(w, r)
		}
	}
}

// ContextCreator accept a request and return a context, used in TemplateRenderHandler.
type ContextCreator func(r *http.Request) interface{}

// TemplateRenderHandler is a shortcut function that generate a http.HandlerFunc.
// The generated func use contextCreator to create context and render the templatePath template file.
// If contextCreator is nil, nil will be used as context.
func TemplateRenderHandler(templatePath string, contextCreator ContextCreator, log *logger.Logger) http.HandlerFunc {
	if log == nil {
		log = l
	}
	return func(w http.ResponseWriter, r *http.Request) {
		defer recover()

		log.Info("Recieve a template render request of", templatePath, "from ip", r.RemoteAddr)

		var err error

		if contextCreator != nil {
			err = RenderTemplate(templatePath, w, contextCreator(r))
		} else {
			err = RenderTemplate(templatePath, w, nil)
		}

		if err != nil {
			log.Warn("Render template", templatePath, "with data", nil, "error: ", err)
		}

		log.Info("Render template", templatePath, "successfully")
	}
}

// RequestFilter accept a http.HandlerFunc and return a new one
// which only accept path is pathMustBe and method is methodMustBe.
// Error message in new hander will be print with logger log, if log is nil, will use default logger.
// If pathMustBe or methodMustBe is empty string, no check will be performed.
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
				log.Warn("Someone visit page", r.URL.Path, "with method", r.Method, ", only", methodMustBe, "is allowed.")
				return
			}
		}

		handlerFunc(w, r)
	}
}
