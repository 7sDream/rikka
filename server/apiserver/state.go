package apiserver

import (
	"net/http"

	"github.com/7sDream/rikka/common/util"
)

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
