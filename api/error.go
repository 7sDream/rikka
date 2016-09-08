package api

// API error message consts
var (
	// Upload errors
	ErrPwdErrMsg         = "Error password"
	InvalidFromArgErrMsg = "From argument can only be website or api"
	NotAImgFileErrMsg    = "The file you upload is not an image"

	// Task errors
	TaskNotExistErrMsg     = "Task not exist"
	TaskAlreadyExistErrMsg = "Task already exist"
	TaskNotFinishErrMsg    = "Task is not finished"
)
