package apiserver

import (
	"net/http"

	"github.com/7sDream/rikka/common/util"
)

func urlHandleFunc(w http.ResponseWriter, r *http.Request) {

	taskID := util.GetTaskIDByRequest(r)

	l.Info("Recieve a url request of task", taskID, "from ip", r.RemoteAddr)

	var jsonData []byte
	var err error

	if jsonData, err = getURLJSON(taskID, r, nil); err != nil {
		l.Error("Error happened when get url json of task", taskID, ":", err)
	} else {
		l.Info("Get url json", string(jsonData), "of task", taskID, "successfully")
	}

	renderJSONOrError(w, taskID, jsonData, err, http.StatusInternalServerError)
}
