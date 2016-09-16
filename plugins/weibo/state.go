package weibo

import (
	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/plugins"
)

const (
	stateUploading     = "uploading"
	stateUploadingCode = 2
	stateUploadingDesc = "Rikka is uploading your image to UPai cloud"
)

func buildUploadingState(taskID string) *api.State {
	return &api.State{
		TaskID:      taskID,
		State:       stateUploading,
		StateCode:   stateUploadingCode,
		Description: stateUploadingDesc,
	}
}

func (wbp weiboPlugin) StateRequestHandler(taskID string) (*api.State, error) {
	l.Debug("Recieve a state request of taskID", taskID)

	pState, err := plugins.GetTaskState(taskID)
	if err != nil {
		l.Error("State of task", taskID, "not found, return error")
		return nil, err
	}

	if pState.StateCode == api.StateErrorCode {
		l.Warn("Get a error state of task", taskID, *pState)
	} else {
		l.Debug("Get a normal state of task", taskID, *pState)
	}
	return pState, nil
}
