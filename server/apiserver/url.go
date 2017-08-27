package apiserver

import (
	"net/http"

	"github.com/7sDream/rikka/common/util"
)

func urlHandleFunc(w http.ResponseWriter, r *http.Request) {
	ip := util.GetClientIP(r)
	taskID := util.GetTaskIDByRequest(r)

	l.Info("Receive a url request of task", taskID, "from ip", ip)

	var jsonData []byte
	var err error

	if jsonData, err = getUrlJson(taskID, r, nil); err != nil {
		l.Error("Error happened when get url json of task", taskID, "request by", ip, ":", err)
	} else {
		l.Debug("Get url json", string(jsonData), "of task", taskID, "request by", ip, "successfully")
	}

	renderJsonOrError(w, taskID, jsonData, err, http.StatusInternalServerError)
}
