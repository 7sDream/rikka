package api

// State shows a state of task.
type State struct {
	TaskID      string
	StateCode   int
	State       string
	Description string
}

// Error struct used to build json from error string.
type Error struct {
	Error string
}

// URL struct used to build json from error URL.
type URL struct {
	URL string
}

// TaskId struct used to build json from taskID.
type TaskId struct {
	TaskId string
}
