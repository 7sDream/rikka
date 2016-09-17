package weibo

import (
	"strconv"
	"sync/atomic"

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

func (wbp weiboPlugin) StateRequestHandle(taskID string) (*api.State, error) {
	l.Debug("Recieve a state request of taskID", taskID)

	taskIDInt, err := strconv.ParseInt(taskID, 10, 64)
	if err != nil {
		l.Fatal("Error happened when parse int from task ID", taskID, ":", err)
	}
	l.Debug("Parse task ID to int successfully")

	pState, err := plugins.GetTaskState(taskID)

	// can't get state
	if err != nil {
		l.Warn("Error happened when get state of task", taskID, ":", err)
		// but task id < counter means task finish
		if taskIDInt <= atomic.LoadInt64(&counter) {
			l.Debug("But task ID <= counter, return finished state")
			return api.BuildFinishState(taskID), nil
		}
		// else, no task fount
		l.Error("task ID > counter, return error")
		return nil, err
	}
	if pState.StateCode == api.StateErrorCode {
		l.Warn("Get a error state of task", taskID, *pState)
	} else {
		l.Debug("Get a normal state of task", taskID, *pState)
	}

	return pState, nil
}
