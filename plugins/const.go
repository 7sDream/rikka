package plugins

// Built in state code and state str, desctiption.
// Must be used in finish and error state.
const (
	StateError     = "error"
	StateErrorCode = -1

	StateFinish            = "finish"
	StateFinishCode        = 0
	StateFinishDescription = "file upload task finish"
)
