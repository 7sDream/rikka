package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/7sDream/rikka/api"
)

func GetState(host string, taskID string) (*api.State, error) {
	url := host + api.StatePath + taskID
	l.Debug("Build state request url:", url)

	res, err := http.Get(url)
	if err != nil {
		l.Debug("Error happened when send state request to url", url, ":", err)
		return nil, err
	}
	l.Debug("Send state request successfully")

	resContent, err := checkRes(url, res)
	if err != nil {
		return nil, err
	}

	pState := &api.State{}
	if err = json.Unmarshal(resContent, pState); err != nil || pState.TaskID == "" {
		l.Debug("Decode response to state json failed, try to decode to error json")
		return nil, mustBeErrorJSON(resContent)
	}
	l.Debug("Decode response to state json successfully")
	return pState, nil
}

func WaitFinish(host string, taskID string) error {
	for {
		state, err := GetState(host, taskID)
		if err != nil {
			l.Debug("Error happened when get state of task", taskID)
			continue
		}

		l.Debug("State of task", taskID, "is:", state.Description)

		if state.StateCode == api.StateErrorCode {
			l.Debug("Rikka return a error task state:", state.Description)
			return errors.New(state.Description)
		}

		if state.StateCode == api.StateFinishCode {
			return nil
		}

		l.Debug("State is not finished, will retry after 1 second...")

		time.Sleep(1 * time.Second)
	}
}
