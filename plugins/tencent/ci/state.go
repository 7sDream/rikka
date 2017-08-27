package ci

import (
	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/plugins"
)

const (
	stateReading     = "reading file"
	stateReadingCode = 2
	stateReadingDesc = "Rikka is reading your image file"

	stateUploading     = "uploading"
	stateUploadingCode = 3
	stateUploadingDesc = "Rikka is uploading your image to Tencent CI"
)

func buildUploadingState(taskID string) *api.State {
	return &api.State{
		TaskID:      taskID,
		State:       stateUploading,
		StateCode:   stateUploadingCode,
		Description: stateUploadingDesc,
	}
}

func buildReadingState(taskID string) *api.State {
	return &api.State{
		TaskID:      taskID,
		State:       stateReading,
		StateCode:   stateReadingCode,
		Description: stateReadingDesc,
	}
}

func (plugin tcciPlugin) StateRequestHandle(taskID string) (*api.State, error) {
	pState, err := plugins.GetTaskState(taskID)
	if err != nil {
		// Not exist as finished
		return api.BuildFinishState(taskID), nil
	}
	return pState, nil
}
