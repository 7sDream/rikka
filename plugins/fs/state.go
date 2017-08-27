package fs

import (
	pathUtil "path/filepath"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

const (
	stateCopying     = "copying"
	stateCopyingCode = 2
	stateCopyingDesc = "Image is being copied to rikka file system"

	stateCreating     = "creating"
	stateCreatingCode = 3
	stateCreatingDesc = "Creating file in fs to store your image"
)

// A shortcut function to build state we need.
func buildCreatingState(taskID string) *api.State {
	return &api.State{
		TaskID:      taskID,
		StateCode:   stateCreatingCode,
		State:       stateCreating,
		Description: stateCreatingDesc,
	}
}

// A shortcut function to build state we need.
func buildCopyingState(taskID string) *api.State {
	return &api.State{
		TaskID:      taskID,
		StateCode:   stateCopyingCode,
		State:       stateCopying,
		Description: stateCopyingDesc,
	}
}

// StateRequestHandle Will be called when receive a get state request.
func (fsp fsPlugin) StateRequestHandle(taskID string) (pState *api.State, err error) {

	l.Debug("Receive a state request of taskID", taskID)

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
	// TaskId not exist or error when get it, check if image file already exist
	if util.IsFile(pathUtil.Join(imageDir, taskID)) {
		// file exist is regarded as a finished state
		pFinishState := api.BuildFinishState(taskID)
		l.Debug("File of task", taskID, "exist, return finished state", *pFinishState)
		return pFinishState, nil
	}

	l.Warn("File of task", taskID, "not exist, get state error:", err)
	// get state error
	return nil, err
}
