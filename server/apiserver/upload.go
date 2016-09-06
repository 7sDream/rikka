package apiserver

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
	"github.com/7sDream/rikka/server/webserver"
)

var (
	taskIDUploading = "[uploading]"
)

// ---- upload handle aux functions --

func checkFromArg(w http.ResponseWriter, r *http.Request) (string, bool) {
	from := r.FormValue("from")
	if from != "website" && from != "api" {
		l.Warn("Someone use a error from value:", from)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("from argument can only be website or api."))
		return "", false
	}
	l.Debug("Request from:", from)
	return from, true
}

func checkPassowrd(w http.ResponseWriter, r *http.Request, from string) bool {
	userPassword := r.FormValue("password")
	if userPassword != password {
		// error password
		l.Warn("Someone input a error password:", userPassword)

		if from == "website" {
			http.Error(w, "Error password", http.StatusUnauthorized)
		} else {
			// from == "api"
			renderErrorJSON(w, taskIDUploading, errors.New(api.ErrPwdErrMsg), http.StatusUnauthorized)
		}

		return false
	}
	l.Debug("Password check successfully")
	return true
}

func checkUploadedFile(w http.ResponseWriter, file multipart.File, from string) bool {
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		l.Error("Error happened when get form file content:", err)

		if from == "website" {
			util.ErrHandle(w, err)
		} else {
			// from == "api"
			renderErrorJSON(w, taskIDUploading, err, http.StatusInternalServerError)
		}

		return false
	}
	l.Debug("Get form file content successfully")

	filetype := http.DetectContentType(fileContent)
	if !strings.HasPrefix(filetype, "image") {
		l.Error("Form file is not a image, it is a", filetype)

		if from == "website" {
			util.ErrHandle(w, errors.New(api.NotAImgFileErrMsg))
		} else {
			// from == "api"
			renderErrorJSON(w, taskIDUploading, errors.New(api.NotAImgFileErrMsg), http.StatusInternalServerError)
		}

		return false
	}
	l.Debug("Check form file type, passed:", filetype)

	if _, err = file.Seek(0, 0); err != nil {
		l.Error("Error when try to seek form file to start:", err)

		if from == "website" {
			util.ErrHandle(w, err)
		} else {
			// from == "api"
			renderErrorJSON(w, taskIDUploading, err, http.StatusInternalServerError)
		}

		return false
	}

	return true
}

func getUploadedFile(w http.ResponseWriter, r *http.Request, from string) (multipart.File, bool) {
	file, _, err := r.FormFile("uploadFile")
	if err != nil {
		// no needed file
		l.Error("Error happened when get form file:", err)

		if from == "website" {
			util.ErrHandle(w, err)
		} else {
			// from == "api"
			renderErrorJSON(w, taskIDUploading, err, http.StatusBadRequest)
		}

		return file, false
	}
	l.Debug("Get uploaded file successfully")

	if !checkUploadedFile(w, file, from) {
		return nil, false
	}

	return file, true
}

func redirectToView(w http.ResponseWriter, taskID string) {
	viewPage := webserver.ViewPath + taskID
	w.Header().Set("Location", viewPage)
	w.WriteHeader(302)
	l.Debug("Redirect user to view page", viewPage)
}

func sendSaveRequestToPlugin(w http.ResponseWriter, file multipart.File, from string) (string, bool) {
	l.Debug("Send file save request to plugin manager")

	pTaskID, err := plugins.AcceptFile(&plugins.SaveRequest{File: file})

	if err != nil {
		l.Error("Error happened when plugin manager process file save request:", err)
		if from == "website" {
			util.ErrHandle(w, err)
		} else {
			renderErrorJSON(w, taskIDUploading, err, http.StatusInternalServerError)
		}
		return "", false
	}

	taskID := pTaskID.TaskID
	l.Debug("Recieve task ID from plugin manager:", taskID)

	return taskID, true
}

func sendUploadResultToClient(w http.ResponseWriter, taskID string, from string) {
	if from == "website" {
		redirectToView(w, taskID)
	} else {
		var taskIDJSON []byte
		var err error
		if taskIDJSON, err = getTaskIDJSON(taskID); err != nil {
			l.Error("Error happened when build task ID json of task", taskID, ":", err)
		} else {
			l.Debug("Build task ID json", taskIDJSON, "of task", taskID, "successfully")
		}
		renderJSONOrError(w, taskID, taskIDJSON, err, http.StatusInternalServerError)
	}
}

// ---- end of upload handle aux functions --

func uploadHandleFunc(w http.ResponseWriter, r *http.Request) {
	defer recover()

	l.Info("Recieve file upload request from ip", r.RemoteAddr)

	maxSize := int64(maxSizeByMB * 1024 * 1024)
	r.Body = http.MaxBytesReader(w, r.Body, maxSize)

	err := r.ParseMultipartForm(maxSize)
	if util.ErrHandle(w, err) {
		l.Error("Error happened when parse form:", err)
		return
	}

	var from string
	var ok bool
	if from, ok = checkFromArg(w, r); !ok {
		return
	}

	if !checkPassowrd(w, r, from) {
		return
	}

	var file multipart.File
	if file, ok = getUploadedFile(w, r, from); !ok {
		return
	}

	var taskID string
	if taskID, ok = sendSaveRequestToPlugin(w, file, from); !ok {
		return
	}

	sendUploadResultToClient(w, taskID, from)
}
