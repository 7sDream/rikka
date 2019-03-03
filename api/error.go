package api

// API error messages
var (
	// Upload errors
	ErrPwdErrMsg         = "error password"
	InvalidFromArgErrMsg = "from argument can only be website or api"
	NotAImgFileErrMsg    = "the file you upload is not an image"

	// Task errors
	TaskNotExistErrMsg     = "task not exist"
	TaskAlreadyExistErrMsg = "task already exist"
	TaskNotFinishErrMsg    = "task is not finished"
)
