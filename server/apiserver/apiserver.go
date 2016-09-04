package apiserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

var password string
var maxSizeByMB float64

var l *logger.Logger

func jsonEncode(obj interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(obj)
	if err == nil {
		return jsonData, nil
	}
	l.Warn("Error happened when encoding", fmt.Sprintf("%+v", obj), "to json :", err)
	return nil, err
}

func getErrorJSON(taskID string, err error) ([]byte, error) {
	obj := plugins.ErrorJSON{
		Error: err.Error(),
	}
	return jsonEncode(obj)
}

func getTaskIDJSON(taskID string) ([]byte, error) {
	obj := plugins.TaskIDJSON{
		TaskID: taskID,
	}
	return jsonEncode(obj)
}

func getStateJSON(taskID string) ([]byte, error) {
	state, err := plugins.GetState(taskID)
	if err != nil {
		l.Warn("Error happened when get state of task", taskID, ":", err)
		return nil, err
	}
	return jsonEncode(state)
}

func getURLJSON(taskID string, r *http.Request, picOp *plugins.ImageOperate) ([]byte, error) {
	url, err := plugins.GetURL(taskID, r, picOp)
	if err != nil {
		l.Warn("Error happened when get url of task", taskID, ":", err)
		return nil, err
	}
	return jsonEncode(url)
}

func stateHandleFunc(w http.ResponseWriter, r *http.Request) {
	defer recover()

	taskID := util.GetTaskIDByRequest(r)

	jsonData, err := getStateJSON(taskID)

	if err != nil {
		// get state json failed
		errorJSONData, err := getErrorJSON(taskID, err)

		if util.ErrHandle(w, err) {
			// get error json failed
			l.Warn("Error happened when get error json of state of task", taskID, ":", err)
			return
		}

		err = util.RenderJSON(w, errorJSONData)

		if util.ErrHandle(w, err) {
			// rander error json failed
			l.Warn("Error happened when render error json of state", errorJSONData, "of task", taskID, ":", err)
			return
		}
	}

	err = util.RenderJSON(w, jsonData)
	if util.ErrHandle(w, err) {
		// render normal json failed
		l.Warn("Error happened when render state json", jsonData, "of task", taskID, ":", err)
	}
}

func urlHandleFunc(w http.ResponseWriter, r *http.Request) {
	defer recover()

	taskID := util.GetTaskIDByRequest(r)

	jsonData, err := getStateJSON(taskID)

	if err != nil {
		// get url json failed
		errorJSONData, err := getURLJSON(taskID, r, nil)

		if util.ErrHandle(w, err) {
			// get error json failed
			l.Warn("Error happened when get error json of url of task", taskID, ":", err)
			return
		}

		err = util.RenderJSON(w, errorJSONData)

		if util.ErrHandle(w, err) {
			// rander error json failed
			l.Warn("Error happened when render error json of url", errorJSONData, "of task", taskID, ":", err)
			return
		}
	}

	err = util.RenderJSON(w, jsonData)
	if util.ErrHandle(w, err) {
		// render normal json failed
		l.Warn("Error happened when render url json", jsonData, "of task", taskID, ":", err)
	}
}

func uploadHandleFunc(w http.ResponseWriter, r *http.Request) {
	defer recover()

	l.Info("Recieve file upload request")

	maxSize := int64(maxSizeByMB * 1024 * 1024)

	r.Body = http.MaxBytesReader(w, r.Body, maxSize)

	err := r.ParseMultipartForm(maxSize)
	if util.ErrHandle(w, err) {
		l.Warn("Error happened when parse form:", err)
		return
	}

	from := r.FormValue("from")
	if from != "website" && from != "api" {
		l.Warn("Someone use a error from value:", from)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("from argument can only be website or api."))
		return
	}

	l.Info("Request from:", from)

	userPassword := r.FormValue("password")
	if userPassword != password {
		// error password
		l.Warn("Someone input a error password:", userPassword)

		if from == "website" {
			http.Error(w, "Error password", http.StatusUnauthorized)
			return
		}

		// from == "api"
		errorJSON, err := getErrorJSON("uploadTask", errors.New("Error password"))
		if util.ErrHandle(w, err) {
			// build error json failed
			l.Warn("Error happened when build error json of url of task", "uploadTask", ":", err)
			return
		}

		// get error json successfully
		err = util.RenderJSON(w, errorJSON)

		// check if json render error
		util.ErrHandle(w, err)
		return
	}

	l.Info("Password check successfully")

	file, _, err := r.FormFile("uploadFile")
	if err != nil {
		// no needed file
		l.Warn("Error happened when get form file:", err)

		if from == "website" {
			util.ErrHandle(w, err)
			return
		}

		// from == "api"
		errorJSON, err := getErrorJSON("uploadTask", err)
		if util.ErrHandle(w, err) {
			// build error json failed
			l.Warn("Error happened when build error json of url of task", "uploadTask", ":", err)
			return
		}

		// get error json successfully
		err = util.RenderJSON(w, errorJSON)

		// check if json render error
		util.ErrHandle(w, err)
		return
	}

	l.Info("Send file save request to plugin")
	taskID, err := plugins.AcceptFile(&plugins.SaveRequest{File: file})
	// return

	if from == "website" {
		// accept file request error
		if util.ErrHandle(w, err) {
			l.Warn("Error happened when plugin process file save request:", err)
			return
		}

		// accept file request successfully
		l.Info("Get taskID:", taskID)
		viewPage := "/view/" + taskID
		w.Header().Set("Location", viewPage)
		w.WriteHeader(302)
		l.Info("Redirect user to view page", viewPage)
		return
	}

	// from == "api"

	// accept file request error
	if err != nil {
		l.Warn("Error happened when plugin process file save request:", err)

		errorJSON, err := getErrorJSON(taskID, err)
		if util.ErrHandle(w, err) {
			// build error json failed
			l.Warn("Error happened when build error json of url of task", taskID, ":", err)
			return
		}

		// build error json successfully
		l.Info("Build error json", errorJSON, "successfully")
		err = util.RenderJSON(w, errorJSON)

		// check if json render successfully
		if util.ErrHandle(w, err) {
			// failed
			l.Warn("Render error json", errorJSON, "failed")
		} else {
			// successfully
			l.Info("Error json", errorJSON, "rendered successfully")
		}
		return
	}

	// accept file request successfully
	l.Info("Get taskID:", taskID)
	taskIDJSON, err := getTaskIDJSON(taskID)

	// get taskID json error
	if err != nil {
		l.Warn("Error happened when get taskID json of task", taskID, ":", err)

		errorJSON, err := getErrorJSON(taskID, err)

		// build error json error
		if util.ErrHandle(w, err) {
			l.Warn("Error happened when build error json of task", taskID, ":", err)
			return
		}

		// build error json successfully
		l.Info("Build error json", errorJSON, "successfully")
		err = util.RenderJSON(w, errorJSON)

		// check if json rend successfully
		if util.ErrHandle(w, err) {
			// failed
			l.Warn("Render error json,", errorJSON, "failed")
		} else {
			// successfully
			l.Info("Error json", errorJSON, "rendered successfully")
		}
		return
	}

	// get taskID json successfully
	l.Info("Build taskID json", taskIDJSON, "successfully")
	err = util.RenderJSON(w, taskIDJSON)

	// check if json render successfully
	if util.ErrHandle(w, err) {
		l.Warn("Render taskID json", taskIDJSON, "error:", err)
	} else {
		l.Info("Render taskID json", taskIDJSON, "successfully")
	}
}

// StartRikkaAPIServer start API server of Rikka
func StartRikkaAPIServer(argPassword string, argMaxSizeByMb float64, log *logger.Logger) {

	password = argPassword
	maxSizeByMB = argMaxSizeByMb

	l = log.SubLogger("[API]")

	stateHandler := util.RequestFilter(
		"", "GET", l,
		stateHandleFunc,
	)

	urlHandler := util.RequestFilter(
		"", "GET", l,
		urlHandleFunc,
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
