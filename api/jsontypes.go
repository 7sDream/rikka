package api

// State shows a state of task.
type State struct {
	TaskID      string
	StateCode   int
	State       string
	Description string
}

// Error struct used to build json from eror string.
type Error struct {
	Error string
}

// URL struct used to build json from eror URL.
type URL struct {
	URL string
}

// TaskID struct used to build json from taskID.
type TaskID struct {
	TaskID string
}
