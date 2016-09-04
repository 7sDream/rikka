package plugins

import (
	"errors"
	"sync"
)

var tasks = struct {
	sync.RWMutex
	m map[string]*State
}{m: make(map[string]*State)}

// CreateTask add a task to task list.
// If taskID already exist, return an error.
func CreateTask(taskID string) error {
	tasks.Lock()
	defer tasks.Unlock()

	if _, ok := tasks.m[taskID]; ok { // key exist
		return errors.New("Task already exist")
	}

	createState := BuildCreateState(taskID)
	tasks.m[taskID] = &createState
	return nil
}

// ChangeTaskState change the state of a task.
// If taskID not exist, return an error.
func ChangeTaskState(state State) error {
	tasks.Lock()
	defer tasks.Unlock()

	if _, ok := tasks.m[state.TaskID]; ok { // key exist
		pState := tasks.m[state.TaskID]
		pState.StateCode = state.StateCode
		pState.State = state.State
		pState.Description = state.Description
		return nil
	}

	return errors.New("Task not exist")
}

// GetTaskState get state of a task.
// If task not exist, return an error.
func GetTaskState(taskID string) (pState *State, err error) {
	tasks.RLock()
	defer tasks.RUnlock()

	if pState, ok := tasks.m[taskID]; ok { // key exist
		return pState, nil
	}
	return nil, errors.New("Task not exist.")
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

	return errors.New("Task not exist")
}

// BuildCreateState build a standard just-create state from taskID.
func BuildCreateState(taskID string) State {
	return State{
		TaskID:      taskID,
		State:       StateCreate,
		StateCode:   StateCreateCode,
		Description: StateCreateDescription,
	}
}

// BuildFinishState build a standard finished state from taskID.
func BuildFinishState(taskID string) State {
	return State{
		TaskID:      taskID,
		State:       StateFinish,
		StateCode:   StateFinishCode,
		Description: StateFinishDescription,
	}
}

// BuildErrorState build a standard error state from taskID and description.
func BuildErrorState(taskID string, description string) State {
	return State{
		TaskID:      taskID,
		State:       StateError,
		StateCode:   StateErrorCode,
		Description: description,
	}
}
