package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/7sDream/rikka/plugins"
)

const stateAPIPath = "/api/state/"

func getState(host string, taskID string) *plugins.State {

	url := host + stateAPIPath + taskID
	l.Debug("Build state request url:", url)

	res, err := http.Get(url)
	if err != nil {
		l.Fatal("Error when send state request to", url, ":", err)
	}
	l.Debug("Send state request successfully")

	resContent := checkRes(url, res)

	stateJSON := &plugins.State{}
	if err = json.Unmarshal(resContent, stateJSON); err == nil {
		if stateJSON.TaskID != "" {
			l.Debug("Decode response to state json")
			return stateJSON
		}
	}
	l.Debug("Decode response to state json failed, try to decode to error message")

	mustBeErrorJSON(resContent)

	return nil
}

func waitFinish(host string, taskID string) {
	for {
		state := getState(host, taskID)

		if state.StateCode == plugins.StateErrorCode {
			l.Fatal("Task state error:", state.Description)
		}

		if state.StateCode == plugins.StateFinishCode {
			return
		}

		l.Info("State if task is:", state.Description)

		time.Sleep(1 * time.Second)
	}
}
