package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

var l *logger.Logger

func jsonEncode(taskID string, what string, obj interface{}, err error) ([]byte, error) {
	if err == nil {
		jsonData, err := json.Marshal(obj)
		if err == nil {
			return jsonData, nil
		}
		l.Error("Error happened when encoding", fmt.Sprintf("%+v", obj), "to json :", err)
	} else {
		l.Error("Error happened when get", what, "of task", taskID, ":", err)
	}
	return nil, err
}

func getStateJSON(taskID string) ([]byte, error) {
	state, err := plugins.GetState(taskID)
	return jsonEncode(taskID, "state", state, err)
}

func stateHandleFunc(w http.ResponseWriter, r *http.Request) {
	defer recover()

	taskID := util.GetTaskIDByRequest(r)

	jsonData, err := getStateJSON(taskID)

	if util.ErrHandle(w, err) {
		return
	}

	err = util.RenderJSON(w, jsonData)
	util.ErrHandle(w, err)
}

func getURLJSON(taskID string, r *http.Request, picOp *plugins.PictureOperate) ([]byte, error) {
	url, err := plugins.GetURL(taskID, r, picOp)
	return jsonEncode(taskID, "url", url, err)
}

func urlHandleFunc(w http.ResponseWriter, r *http.Request) {
	defer recover()

	taskID := util.GetTaskIDByRequest(r)

	jsonData, err := getURLJSON(taskID, r, nil)

	if util.ErrHandle(w, err) {
		return
	}

	err = util.RenderJSON(w, jsonData)
	util.ErrHandle(w, err)
}

// StartRikkaAPIServer start API server of Rikka
func StartRikkaAPIServer(log *logger.Logger) {
	l = log.SubLogger("[API]")

	stateHandler := util.RequestFilter(
		"", "GET", l,
		stateHandleFunc,
	)

	urlHandler := util.RequestFilter(
		"", "GET", l,
		urlHandleFunc,
	)

	http.HandleFunc("/api/state/", stateHandler)
	http.HandleFunc("/api/url/", urlHandler)

	l.Info("API server start successfully")
	return
}
