package cos

import (
	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/plugins"
)

const (
	stateUploading     = "uploading"
	stateUploadingCode = 2
	stateUploadingDesc = "Rikka is uploading your image to Tencent COS"
)

func buildUploadingState(taskID string) *api.State {
	return &api.State{
		TaskID:      taskID,
		State:       stateUploading,
		StateCode:   stateUploadingCode,
		Description: stateUploadingDesc,
	}
}

func (cosp tccosPlugin) StateRequestHandle(taskID string) (*api.State, error) {
	pState, err := plugins.GetTaskState(taskID)
	if err != nil {
		// Not exist as finished
		return api.BuildFinishState(taskID), nil
	}
	return pState, nil
}
