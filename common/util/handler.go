package util

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net"
	"net/http"
	pathutil "path/filepath"
	"strings"

	"github.com/7sDream/rikka/common/logger"
)

// GetTaskIDByRequest gets last part of url path as a taskID and return it.
func GetTaskIDByRequest(r *http.Request) string {
	splitedPath := strings.Split(r.URL.Path, "/")
	filename := splitedPath[len(splitedPath)-1]
	return filename
}

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

// GetClientIP get client ip address from a http request.
// Try to get ip from x-forwarded-for header first,
// If key not exist in header, try to get ip:host from r.RemoteAddr
// split it to ip:port and return ip, If error happened, return 0.0.0.0
func GetClientIP(r *http.Request) string {
	defer func() {
		if r := recover(); r != nil {
			l.Error("Unexcepted panic happened when get client ip:", r)
		}
	}()

	forwardIP := r.Header.Get("X-FORWARDED-FOR")
	if forwardIP != "" {
		return forwardIP
	}

	socket := r.RemoteAddr
	host, _, err := net.SplitHostPort(socket)
	if err != nil {
		l.Warn("Error happened when get IP address :", err)
		return "0.0.0.0"
	}
	return host
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

// DisableListDir accept a handle func and return a handle that not allow list dir.
func DisableListDir(log *logger.Logger, h http.HandlerFunc) http.HandlerFunc {
	if log == nil {
		l.Warn("Get a nil logger in function DisableListDirFunc")
		log = l
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			log.Warn(GetClientIP(r), "try to list dir", r.URL.Path)
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
		l.Warn("Get a nil logger in function TemplateRenderHandler")
		log = l
	}
	return func(w http.ResponseWriter, r *http.Request) {

		templateName := pathutil.Base(templatePath)
		ip := GetClientIP(r)

		log.Info("Recieve a template render request of", templateName, "from ip", ip)

		var data interface{}
		if contextCreator != nil {
			data = contextCreator(r)
		} else {
			data = nil
		}

		err := RenderTemplate(templatePath, w, data)
		if err != nil {
			log.Warn("Error happened when render template", templateName, "with data", fmt.Sprintf("%+v", data), "to", ip, ": ", err)
		}

		log.Info("Render template", templateName, "to", ip, "successfully")
	}
}

// RequestFilter accept a http.HandlerFunc and return a new one
// which only accept path is pathMustBe and method is methodMustBe.
// Error message in new hander will be print with logger log, if log is nil, will use default logger.
// If pathMustBe or methodMustBe is empty string, no check will be performed.
func RequestFilter(pathMustBe string, methodMustBe string, log *logger.Logger, handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if log == nil {
			l.Warn("Get a nil logger in function RequestFilter")
			log = l
		}

		ip := GetClientIP(r)

		if pathMustBe != "" {
			if !MustBeOr404(w, r, pathMustBe) {
				log.Warn(ip, "visit a non-exist page", r.URL.Path, ", excepted is /")
				return
			}
		}

		if methodMustBe != "" {
			if !CheckMethod(w, r, methodMustBe) {
				log.Warn(ip, "visit page", r.URL.Path, "with method", r.Method, ", only", methodMustBe, "is allowed")
				return
			}
		}

		handlerFunc(w, r)
	}
}
