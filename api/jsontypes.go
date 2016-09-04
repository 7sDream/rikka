package api

// StateJSON shows a state of task.
type State struct {
	TaskID      string
	StateCode   int
	State       string
	Description string
}

// ErrorJSON struct used to build json from eror string.
type Error struct {
	Error string
}

// URLJSON struct used to build json from eror URL.
type URL struct {
	URL string
}

// TaskIDJSON struct used to build json from taskID.
type TaskID struct {
	TaskID string
}
