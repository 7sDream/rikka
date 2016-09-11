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
)

var (
	viewPath        = ""
	taskIDUploading = "[uploading]"
	acceptedTypes   = []string{
		"jpeg", "bmp", "gif", "png",
	}
)

// ---- upload handle aux functions --

func checkFromArg(w http.ResponseWriter, r *http.Request, ip string) (string, bool) {
	from := r.FormValue("from")
	if from != "website" && from != "api" {
		l.Warn(ip, "use a error from value:", from)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(api.InvalidFromArgErrMsg))
		return "", false
	}
	l.Debug("Request of", ip, "is from:", from)
	return from, true
}

func checkPassowrd(w http.ResponseWriter, r *http.Request, ip string, from string) bool {
	userPassword := r.FormValue("password")
	if userPassword != password {
		// error password
		l.Warn(ip, "input a error password:", userPassword)

		if from == "website" {
			http.Error(w, "Error password", http.StatusUnauthorized)
		} else {
			// from == "api"
			renderErrorJSON(w, taskIDUploading, errors.New(api.ErrPwdErrMsg), http.StatusUnauthorized)
		}

		return false
	}
	l.Debug("Password check for", ip, "successfully")
	return true
}

// IsAccepted check a mime filetype is accped by rikka.
func IsAccepted(fileMimeTypeStr string) bool {
	if !strings.HasPrefix(fileMimeTypeStr, "image") {
		return false
	}
	for _, acceptedType := range acceptedTypes {
		if strings.HasSuffix(fileMimeTypeStr, "/"+acceptedType) {
			return true
		}
	}
	return false
}

func checkUploadedFile(w http.ResponseWriter, file multipart.File, ip string, from string) (bool, int64) {
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		l.Error("Error happened when get form file content of ip", ip, ":", err)

		if from == "website" {
			util.ErrHandle(w, err)
		} else {
			// from == "api"
			renderErrorJSON(w, taskIDUploading, err, http.StatusInternalServerError)
		}

		return false, 0
	}
	l.Debug("Get form file content of ip", ip, "successfully")

	filetype := http.DetectContentType(fileContent)

	if !IsAccepted(filetype) {
		l.Error("Form file submitted by", ip, "is not a image, it is a", filetype)

		if from == "website" {
			util.ErrHandle(w, errors.New(api.NotAImgFileErrMsg))
		} else {
			// from == "api"
			renderErrorJSON(w, taskIDUploading, errors.New(api.NotAImgFileErrMsg), http.StatusInternalServerError)
		}

		return false, 0
	}
	l.Debug("Check type of form file submitted by", ip, ", passed:", filetype)

	if _, err = file.Seek(0, 0); err != nil {
		l.Error("Error when try to seek form file submitted by", ip, "to start:", err)

		if from == "website" {
			util.ErrHandle(w, err)
		} else {
			// from == "api"
			renderErrorJSON(w, taskIDUploading, err, http.StatusInternalServerError)
		}

		return false, 0
	}

	return true, int64(len((fileContent)))
}

func getUploadedFile(w http.ResponseWriter, r *http.Request, ip string, from string) (multipart.File, int64, bool) {
	file, _, err := r.FormFile("uploadFile")
	if err != nil {
		// no needed file
		l.Error("Error happened when get form file from request of", ip, ":", err)

		if from == "website" {
			util.ErrHandle(w, err)
		} else {
			// from == "api"
			renderErrorJSON(w, taskIDUploading, err, http.StatusBadRequest)
		}

		return nil, 0, false
	}
	l.Debug("Get uploaded file from request of", ip, "successfully")

	acceptable, fileSize := checkUploadedFile(w, file, ip, from)
	if !acceptable {
		return nil, 0, false
	}

	return file, fileSize, true
}

func redirectToView(w http.ResponseWriter, r *http.Request, ip string, taskID string) {
	viewPage := viewPath + taskID
	http.Redirect(w, r, viewPage, http.StatusFound)
	l.Debug("Redirect client", ip, "to view page", viewPage)
}

func sendSaveRequestToPlugin(w http.ResponseWriter, file multipart.File, fileSize int64, ip string, from string) (string, bool) {
	l.Debug("Send file save request to plugin manager")

	pTaskID, err := plugins.AcceptFile(&plugins.SaveRequest{
		File:     file,
		FileSize: fileSize,
	})

	if err != nil {
		l.Error("Error happened when plugin manager process file save request by ip", ip, ":", err)
		if from == "website" {
			util.ErrHandle(w, err)
		} else {
			renderErrorJSON(w, taskIDUploading, err, http.StatusInternalServerError)
		}
		return "", false
	}

	taskID := pTaskID.TaskID
	l.Info("Recieve task ID request by", ip, "from plugin manager:", taskID)

	return taskID, true
}

func sendUploadResultToClient(w http.ResponseWriter, r *http.Request, ip string, taskID string, from string) {
	if from == "website" {
		redirectToView(w, r, ip, taskID)
	} else {
		var taskIDJSON []byte
		var err error
		if taskIDJSON, err = getTaskIDJSON(taskID); err != nil {
			l.Error("Error happened when build task ID json of task", taskID, "request by", ip, ":", err)
		} else {
			l.Info("Build task ID json", taskIDJSON, "of task", "request by", ip, "successfully")
		}
		renderJSONOrError(w, taskID, taskIDJSON, err, http.StatusInternalServerError)
	}
}

// ---- end of upload handle aux functions --

func uploadHandleFunc(w http.ResponseWriter, r *http.Request) {
	ip := util.GetClientIP(r)

	l.Info("Recieve file upload request from ip", ip)

	maxSize := int64(maxSizeByMb * 1024 * 1024)
	r.Body = http.MaxBytesReader(w, r.Body, maxSize)

	err := r.ParseMultipartForm(maxSize)
	if util.ErrHandle(w, err) {
		l.Error("Error happened when parse form submitted by", ip, ":", err)
		return
	}

	var from string
	var ok bool
	if from, ok = checkFromArg(w, r, ip); !ok {
		return
	}

	if !checkPassowrd(w, r, ip, from) {
		return
	}

	var file multipart.File
	var fileSize int64
	if file, fileSize, ok = getUploadedFile(w, r, ip, from); !ok {
		return
	}

	var taskID string
	if taskID, ok = sendSaveRequestToPlugin(w, file, fileSize, ip, from); !ok {
		return
	}

	sendUploadResultToClient(w, r, ip, taskID, from)
}
