package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/7sDream/rikka/api"
)

func getState(host string, taskID string) *api.State {

	url := host + api.StatePath + taskID
	l.Debug("Build state request url:", url)

	res, err := http.Get(url)
	if err != nil {
		l.Fatal("Error when send state request to", url, ":", err)
	}
	l.Debug("Send state request successfully")

	resContent := checkRes(url, res)

	pError := &api.State{}
	if err = json.Unmarshal(resContent, pError); err == nil {
		if pError.TaskID != "" {
			l.Debug("Decode response to state json")
			return pError
		}
	}
	l.Debug("Decode response to state json failed, try to decode to error message")

	mustBeErrorJSON(resContent)

	return nil
}

func waitFinish(host string, taskID string) {
	for {
		state := getState(host, taskID)

		l.Info("State of task", taskID, "is:", state.Description)

		if state.StateCode == api.StateErrorCode {
			l.Fatal("Task state error:", state.Description)
		}

		if state.StateCode == api.StateFinishCode {
			return
		}

		l.Warn("State is not finished, will retry after 1 second...")

		time.Sleep(1 * time.Second)
	}
}
