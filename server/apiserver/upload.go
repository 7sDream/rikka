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
	from := r.FormValue(api.FormKeyFrom)
	if from != api.FromWebsite && from != api.FromAPI {
		l.Warn(ip, "use a error from value:", from)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(api.InvalidFromArgErrMsg))
		return "", false
	}
	l.Debug("Request of", ip, "is from:", from)
	return from, true
}

func checkPassword(w http.ResponseWriter, r *http.Request, ip string, from string) bool {
	userPassword := r.FormValue(api.FormKeyPWD)
	if userPassword != password {
		// error password
		l.Warn(ip, "input a error password:", userPassword)

		if from == api.FromWebsite {
			http.Error(w, "Error password", http.StatusUnauthorized)
		} else {
			// from == "api"
			renderErrorJson(w, taskIDUploading, errors.New(api.ErrPwdErrMsg), http.StatusUnauthorized)
		}

		return false
	}
	l.Debug("Password check for", ip, "successfully")
	return true
}

// IsAccepted check a mime file type is accepted by rikka.
func IsAccepted(fileMimeTypeStr string) (string, bool) {
	if !strings.HasPrefix(fileMimeTypeStr, "image") {
		return "", false
	}
	for _, acceptedType := range acceptedTypes {
		if strings.HasSuffix(fileMimeTypeStr, "/"+acceptedType) {
			return acceptedType, true
		}
	}
	return "", false
}

func checkUploadedFile(w http.ResponseWriter, file multipart.File, ip string, from string) (*plugins.SaveRequest, bool) {
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		l.Error("Error happened when get form file content of ip", ip, ":", err)

		if from == api.FromWebsite {
			util.ErrHandle(w, err)
		} else {
			// from == "api"
			renderErrorJson(w, taskIDUploading, err, http.StatusInternalServerError)
		}

		return nil, false
	}
	l.Debug("Get form file content of ip", ip, "successfully")

	fileType := http.DetectContentType(fileContent)

	ext, ok := IsAccepted(fileType)

	if !ok {
		l.Error("Form file submitted by", ip, "is not a image, it is a", fileType)

		if from == api.FromWebsite {
			util.ErrHandle(w, errors.New(api.NotAImgFileErrMsg))
		} else {
			// from == "api"
			renderErrorJson(w, taskIDUploading, errors.New(api.NotAImgFileErrMsg), http.StatusInternalServerError)
		}

		return nil, false
	}
	l.Debug("Check type of form file submitted by", ip, ", passed:", fileType)

	if _, err = file.Seek(0, 0); err != nil {
		l.Error("Error when try to seek form file submitted by", ip, "to start:", err)

		if from == api.FromWebsite {
			util.ErrHandle(w, err)
		} else {
			// from == "api"
			renderErrorJson(w, taskIDUploading, err, http.StatusInternalServerError)
		}

		return nil, false
	}

	return &plugins.SaveRequest{
		File:     file,
		FileSize: int64(len(fileContent)),
		FileExt:  ext,
	}, true
}

func getUploadedFile(w http.ResponseWriter, r *http.Request, ip string, from string) (*plugins.SaveRequest, bool) {
	file, _, err := r.FormFile(api.FormKeyFile)
	if err != nil {
		// no needed file
		l.Error("Error happened when get form file from request of", ip, ":", err)

		if from == api.FromWebsite {
			util.ErrHandle(w, err)
		} else {
			// from == "api"
			renderErrorJson(w, taskIDUploading, err, http.StatusBadRequest)
		}

		return nil, false
	}
	l.Debug("Get uploaded file from request of", ip, "successfully")

	pSaveRequest, ok := checkUploadedFile(w, file, ip, from)
	if !ok {
		return nil, false
	}

	return pSaveRequest, true
}

func redirectToView(w http.ResponseWriter, r *http.Request, ip string, taskID string) {
	viewPage := viewPath + taskID
	http.Redirect(w, r, viewPage, http.StatusFound)
	l.Debug("Redirect client", ip, "to view page", viewPage)
}

func sendSaveRequestToPlugin(w http.ResponseWriter, pStateRequest *plugins.SaveRequest, ip string, from string) (string, bool) {
	l.Debug("Send file save request to plugin manager")

	pTaskID, err := plugins.AcceptFile(pStateRequest)

	if err != nil {
		l.Error("Error happened when plugin manager process file save request by ip", ip, ":", err)
		if from == api.FromWebsite {
			util.ErrHandle(w, err)
		} else {
			renderErrorJson(w, taskIDUploading, err, http.StatusInternalServerError)
		}
		return "", false
	}

	taskID := pTaskID.TaskId
	l.Info("Receive task ID request by", ip, "from plugin manager:", taskID)

	return taskID, true
}

func sendUploadResultToClient(w http.ResponseWriter, r *http.Request, ip string, taskID string, from string) {
	if from == api.FromWebsite {
		redirectToView(w, r, ip, taskID)
	} else {
		var taskIDJSON []byte
		var err error
		if taskIDJSON, err = getTaskIdJson(taskID); err != nil {
			l.Error("Error happened when build task ID json of task", taskID, "request by", ip, ":", err)
		} else {
			l.Info("Build task ID json", taskIDJSON, "of task", "request by", ip, "successfully")
		}
		renderJsonOrError(w, taskID, taskIDJSON, err, http.StatusInternalServerError)
	}
}

// ---- end of upload handle aux functions --

func uploadHandleFunc(w http.ResponseWriter, r *http.Request) {
	ip := util.GetClientIP(r)

	l.Info("Receive file upload request from ip", ip)

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

	if !checkPassword(w, r, ip, from) {
		return
	}

	var pSaveRequest *plugins.SaveRequest
	if pSaveRequest, ok = getUploadedFile(w, r, ip, from); !ok {
		return
	}

	var taskID string
	if taskID, ok = sendSaveRequestToPlugin(w, pSaveRequest, ip, from); !ok {
		return
	}

	sendUploadResultToClient(w, r, ip, taskID, from)
}
