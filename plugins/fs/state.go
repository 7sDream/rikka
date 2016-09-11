package fs

import (
	pathutil "path/filepath"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

const (
	stateCopying     = "copying"
	stateCopyingCode = 2
	stateCopyingDesc = "Image is being copied to rikka file system"
)

// A shortcut funtion to build state we need.
func buildCopyingState(taskID string) *api.State {
	return &api.State{
		TaskID:      taskID,
		StateCode:   stateCopyingCode,
		State:       stateCopying,
		Description: stateCopyingDesc,
	}
}

// StateRequestHandle Will be called when recieve a get state request.
func (fsp fsPlugin) StateRequestHandle(taskID string) (pState *api.State, err error) {

	l.Debug("Recieve a state request of taskID", taskID)

	// taskID exist on task list, just return it
	if pState, err = plugins.GetTaskState(taskID); err == nil {
		if pState.StateCode == api.StateErrorCode {
			l.Warn("Get a error state of task", taskID, *pState)
		} else {
			l.Debug("Get a normal state of task", taskID, *pState)
		}
		return pState, nil
	}

	l.Debug("State of task", taskID, "not found, check if file exist")
	// TaskID not exist or error when get it, check if image file already exist
	if util.IsFile(pathutil.Join(imageDir, taskID)) {
		// file exist is regarded as a finished state
		pFinishState := api.BuildFinishState(taskID)
		l.Debug("File of task", taskID, "exist, return finished state", *pFinishState)
		return pFinishState, nil
	}

	l.Warn("File of task", taskID, "not exist, get state error:", err)
	// get state error
	return nil, err
}
