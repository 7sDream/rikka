package plugins

import (
	"errors"
	"sync"

	"github.com/7sDream/rikka/api"
)

var (
	tasks = struct {
		sync.RWMutex
		m map[string]*api.State
	}{
		m: make(map[string]*api.State),
	}
)

// CreateTask add a task to task list.
// If taskID already exist, return an error.
func CreateTask(taskID string) error {
	tasks.Lock()
	defer tasks.Unlock()

	if _, ok := tasks.m[taskID]; ok { // key exist
		return errors.New(api.TaskAlreadyExist)
	}

	tasks.m[taskID] = api.BuildCreateState(taskID)
	return nil
}

// ChangeTaskState change the state of a task.
// If taskID not exist, return an error.
func ChangeTaskState(pProvidedState *api.State) error {
	tasks.Lock()
	defer tasks.Unlock()

	if pState, ok := tasks.m[pProvidedState.TaskID]; ok { // key exist
		pState.StateCode = pProvidedState.StateCode
		pState.State = pProvidedState.State
		pState.Description = pProvidedState.Description
		return nil
	}

	return errors.New(api.TaskNotExistErrMsg)
}

// GetTaskState get state of a task.
// If task not exist, return an error.
func GetTaskState(taskID string) (*api.State, error) {
	tasks.RLock()
	defer tasks.RUnlock()

	if pState, ok := tasks.m[taskID]; ok { // key exist
		return pState, nil
	}
	return nil, errors.New(api.TaskNotExistErrMsg)
}

// DeleteTask delete a task from task list.
// If taskID not exist, return an error.
func DeleteTask(taskID string) error {
	tasks.Lock()
	defer tasks.Unlock()

	if _, ok := tasks.m[taskID]; ok { // key exist
		delete(tasks.m, taskID)
		return nil
	}

	return errors.New(api.TaskNotExistErrMsg)
}
