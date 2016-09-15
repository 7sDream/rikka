package qiniu

import (
	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/plugins"
)

const (
	statePreparing     = "preparing"
	statePreparingCode = 2
	statePreparingDesc = "Rikka is preparing to upload your image to Qiniu cloud"

	stateUploading     = "uploading"
	stateUploadingCode = 3
	stateUploadingDesc = "Rikka is uploading your image to Qiniu cloud"
)

func buildPreparingState(taskID string) *api.State {
	return &api.State{
		TaskID:      taskID,
		State:       statePreparing,
		StateCode:   statePreparingCode,
		Description: statePreparingDesc,
	}
}

func buildUploadingState(taskID string) *api.State {
	return &api.State{
		TaskID:      taskID,
		State:       stateUploading,
		StateCode:   stateUploadingCode,
		Description: stateUploadingDesc,
	}
}

func (qnp qiniuPlugin) StateRequestHandle(taskID string) (pState *api.State, err error) {
	l.Debug("Recieve a state request of taskID", taskID)

	pState, err = plugins.GetTaskState(taskID)
	if err == nil {
		if pState.StateCode == api.StateErrorCode {
			l.Warn("Get a error state of task", taskID, *pState)
		} else {
			l.Debug("Get a normal state of task", taskID, *pState)
		}
		return pState, nil
	}

	l.Debug("State of task", taskID, "not found, just return a finish state")
	return api.BuildFinishState(taskID), nil
}
