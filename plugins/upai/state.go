package upai

import "github.com/7sDream/rikka/api"

const (
	stateUploading     = "uploading"
	stateUploadingCode = 3
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
