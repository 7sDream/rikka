package plugins

import "errors"

// CreateTask add a task to task list.
// If taskID already exist, return an error.
func CreateTask(taskID string, stateCode int, state string, desc string) error {
	tasks.Lock()
	defer tasks.Unlock()

	if _, ok := tasks.m[taskID]; ok { // key exist
		return errors.New("Task already exist")
	}

	stateStuct := &State{
		TaskID: taskID, StateCode: stateCode, State: state, Description: desc,
	}
	tasks.m[taskID] = stateStuct
	return nil
}

// ChangeTaskState change the state of a task.
// If taskID not exist, return an error.
func ChangeTaskState(taskID string, stateCode int, state string, desc string) error {
	tasks.Lock()
	defer tasks.Unlock()

	if _, ok := tasks.m[taskID]; ok { // key exist
		pState := tasks.m[taskID]
		pState.StateCode = stateCode
		pState.State = state
		pState.Description = desc
		return nil
	}

	return errors.New("Task not exist")
}

// GetTaskState get state of a task.
// If task not exist, return an error.
func GetTaskState(taskID string) (*State, error) {
	tasks.RLock()
	defer tasks.RUnlock()

	if _, ok := tasks.m[taskID]; ok { // key exist
		return tasks.m[taskID], nil
	}

	return nil, errors.New("Task not exist.")
}

// FinishTask finish a task and delete it from task list.
// If taskID not exist, return an error.
func FinishTask(taskID string) error {
	tasks.Lock()
	defer tasks.Unlock()

	if _, ok := tasks.m[taskID]; ok { // key exist
		delete(tasks.m, taskID)
		return nil
	}

	return errors.New("Task not exist")
}
