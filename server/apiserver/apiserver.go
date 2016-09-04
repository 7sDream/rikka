package apiserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

var password string
var maxSizeByMB float64

var l *logger.Logger

// jsonEncode encode a object to json bytes.
func jsonEncode(obj interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(obj)
	if err == nil {
		l.Debug("Encode data", fmt.Sprint(obj), "to json", string(jsonData), "successfully")
		return jsonData, nil
	}
	l.Error("Error happened when encoding", fmt.Sprint(obj), "to json :", err)
	return nil, err
}

// getErrorJSON get error json bytes like {"Error": "error message"}
func getErrorJSON(err error) ([]byte, error) {
	obj := plugins.ErrorJSON{
		Error: err.Error(),
	}
	return jsonEncode(obj)
}

// getErrorJSON get error json bytes like {"TaskID": "12312398374237"}
func getTaskIDJSON(taskID string) ([]byte, error) {
	obj := plugins.TaskIDJSON{
		TaskID: taskID,
	}
	return jsonEncode(obj)
}

// getStateJSON get state json bytes.
// Will call plgins.GetState
func getStateJSON(taskID string) ([]byte, error) {
	l.Debug("Send state request of task", taskID, "to plugin manager")
	state, err := plugins.GetState(taskID)
	if err != nil {
		l.Warn("Error happened when get state of task", taskID, ":", err)
		return nil, err
	}
	l.Debug("Get state of task", taskID, "successfully")
	return jsonEncode(state)
}

// getURLJSON get url json bytes like {"URL": "http://127.0.0.1/files/filename"}
// Will call plgins.GetURL
func getURLJSON(taskID string, r *http.Request, picOp *plugins.ImageOperate) ([]byte, error) {
	l.Debug("Send url request of task", taskID, "to plugin manager")
	url, err := plugins.GetURL(taskID, r, picOp)
	if err != nil {
		l.Error("Error happened when get url of task", taskID, ":", err)
		return nil, err
	}
	l.Debug("Get url of task", taskID, "successfully")
	return jsonEncode(url)
}

func renderErrorJSON(w http.ResponseWriter, taskID string, err error, errorCode int) {
	errorJSONData, err := getErrorJSON(err)

	if util.ErrHandle(w, err) {
		// build error json failed
		l.Error("Error happened when build error json of task", taskID, ":", err)
		return
	}

	// build error json successfully
	l.Debug("Build error json successfully of task", taskID)
	err = util.RenderJSON(w, errorJSONData, errorCode)

	if util.ErrHandle(w, err) {
		// rander error json failed
		l.Error("Error happened when render error json", errorJSONData, "of task", taskID, ":", err)
	} else {
		l.Debug("Render error json successfully")
	}
}

func renderJSONOrError(w http.ResponseWriter, taskID string, jsonData []byte, err error, errorCode int) {
	// has error
	if err != nil {
		renderErrorJSON(w, taskID, err, errorCode)
	}

	// no error, render json
	err = util.RenderJSON(w, jsonData, http.StatusOK)

	// render json failed
	if util.ErrHandle(w, err) {
		l.Error("Error happened when render json", fmt.Sprint(jsonData), "of task", taskID, ":", err)
	} else {
		l.Debug("Render json", string(jsonData), "of task", taskID, "successfully")
	}
}

// stateHandleFunc is the base handle func of path /api/state/taskID
func stateHandleFunc(w http.ResponseWriter, r *http.Request) {
	defer recover()

	taskID := util.GetTaskIDByRequest(r)

	l.Debug("Recieve a state request of task", taskID)

	var jsonData []byte
	var err error
	if jsonData, err = getStateJSON(taskID); err != nil {
		l.Warn("Error happened when get state json of task", taskID, ":", err)
	} else {
		l.Debug("Get state json of task", taskID, "successfully")
	}

	renderJSONOrError(w, taskID, jsonData, err, http.StatusInternalServerError)
}

func urlHandleFunc(w http.ResponseWriter, r *http.Request) {
	defer recover()

	taskID := util.GetTaskIDByRequest(r)

	l.Debug("Recieve a url request of task", taskID)

	var jsonData []byte
	var err error

	if jsonData, err = getURLJSON(taskID, r, nil); err != nil {
		l.Error("Error happened when get url json of task", taskID, ":", err)
	} else {
		l.Info("Get url json of task", taskID, "successfully")
	}

	renderJSONOrError(w, taskID, jsonData, err, http.StatusInternalServerError)
}

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
			return false
		}

		// from == "api"
		renderErrorJSON(w, "[upload task]", errors.New("Error password"), http.StatusUnauthorized)
		return false
	}
	l.Debug("Password check successfully")
	return true
}

func getUploadedFile(w http.ResponseWriter, r *http.Request, from string) (multipart.File, bool) {
	file, _, err := r.FormFile("uploadFile")
	if err != nil {
		// no needed file
		l.Error("Error happened when get form file:", err)

		if from == "website" {
			util.ErrHandle(w, err)
			return file, false
		}

		// from == "api"
		renderErrorJSON(w, "[upload task]", err, http.StatusBadRequest)
		return file, false
	}
	l.Debug("Get uploaded file successfully")
	return file, true
}

func redirectToView(w http.ResponseWriter, taskID string) {
	viewPage := "/view/" + taskID
	w.Header().Set("Location", viewPage)
	w.WriteHeader(302)
	l.Debug("Redirect user to view page", viewPage)
}

func sendSaveRequestToPlugin(w http.ResponseWriter, file multipart.File, from string) (string, bool) {
	l.Debug("Send file save request to plugin manager")
	taskID, err := plugins.AcceptFile(&plugins.SaveRequest{File: file})

	if err != nil {
		l.Error("Error happened when plugin manager process file save request:", err)
		if from == "website" {
			util.ErrHandle(w, err)
		} else {
			renderErrorJSON(w, taskID, err, http.StatusInternalServerError)
		}
		return taskID, false
	}

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

	l.Debug("Recieve file upload request")

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

// StartRikkaAPIServer start API server of Rikka
func StartRikkaAPIServer(argPassword string, argMaxSizeByMb float64, log *logger.Logger) {

	password = argPassword
	maxSizeByMB = argMaxSizeByMb

	l = log.SubLogger("[API]")

	stateHandler := util.RequestFilter(
		"", "GET", l,
		util.DisableListDirFunc(stateHandleFunc),
	)

	urlHandler := util.RequestFilter(
		"", "GET", l,
		util.DisableListDirFunc(urlHandleFunc),
	)

	uploadHandler := util.RequestFilter(
		"/api/upload", "POST", l,
		uploadHandleFunc,
	)

	http.HandleFunc("/api/state/", stateHandler)
	http.HandleFunc("/api/url/", urlHandler)
	http.HandleFunc("/api/upload", uploadHandler)

	l.Info("API server start successfully")

	return
}
