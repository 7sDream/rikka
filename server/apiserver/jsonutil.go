package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

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
	obj := api.Error{
		Error: err.Error(),
	}
	return jsonEncode(obj)
}

// getErrorJSON get error json bytes like {"TaskID": "12312398374237"}
func getTaskIDJSON(taskID string) ([]byte, error) {
	obj := api.TaskID{
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
