package api

// Built-in state code and state str, description.
// Must be used in finish and error state.
const (
	StateError     = "error"
	StateErrorCode = -1

	StateFinish            = "finish"
	StateFinishCode        = 0
	StateFinishDescription = "file upload task finish"

	StateCreate            = "just created"
	StateCreateCode        = 1
	StateCreateDescription = "the task is created just now, waiting for next operate"
)

// BuildCreateState build a standard just-create state from taskID.
func BuildCreateState(taskID string) *State {
	return &State{
		TaskID:      taskID,
		State:       StateCreate,
		StateCode:   StateCreateCode,
		Description: StateCreateDescription,
	}
}

// BuildFinishState build a standard finished state from taskID.
func BuildFinishState(taskID string) *State {
	return &State{
		TaskID:      taskID,
		State:       StateFinish,
		StateCode:   StateFinishCode,
		Description: StateFinishDescription,
	}
}

// BuildErrorState build a standard error state from taskID and description.
func BuildErrorState(taskID string, description string) *State {
	return &State{
		TaskID:      taskID,
		State:       StateError,
		StateCode:   StateErrorCode,
		Description: description,
	}
}
