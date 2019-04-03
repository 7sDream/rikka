package cos

import (
	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/plugins"
	"github.com/satori/go.uuid"
)

func uploadToCos(q *plugins.SaveRequest, taskID string) {
	defer func() {
		if err := recover(); err != nil {
			l.Error("Panic happened in background:", err)
			var errorMsg string
			switch t := err.(type) {
			case string:
				errorMsg = t
			case error:
				errorMsg = t.Error()
			default:
				errorMsg = "Unknown"
			}
			if err = plugins.ChangeTaskState(api.BuildErrorState(taskID, errorMsg)); err != nil {
				l.Fatal("Unable to change task", taskID, "to error state:", err)
			}
		}
	}()

	err := plugins.ChangeTaskState(buildUploadingState(taskID))
	if err != nil {
		l.Fatal("Error happened when change state of task", taskID, "to uploading:", err)
	}
	l.Debug("Change state of task", taskID, "to uploading successfully")

	l.Debug("Uploading", taskID, "to your cos bucket...")

	if err := client.Upload(q, taskID); err != nil {
		l.Error("Error happened when upload image", taskID, "to cos:", err)
		if err = plugins.ChangeTaskState(api.BuildErrorState(taskID, err.Error())); err != nil {
			l.Fatal("Unable to change task", taskID, "to error state:", err)
		}
		return
	}
	l.Debug("Upload image", taskID, "to cos successfully")

	if err := plugins.DeleteTask(taskID); err != nil {
		l.Error("Error happened when delete task", taskID, ":", err)
	}
	l.Debug("Delete task", taskID, "successfully")
}

func (plugin tencentCloudObjectStoragePlugin) SaveRequestHandle(q *plugins.SaveRequest) (*api.TaskId, error) {
	l.Debug("Receive a file save request")
	taskID := uuid.NewV4().String() + "." + q.FileExt

	err := plugins.CreateTask(taskID)
	if err != nil {
		l.Fatal("Error happened when create new task!")
	}
	l.Debug("create task", taskID, "successfully, starting background task")

	go uploadToCos(q, taskID)

	l.Debug("Background task started, return task ID:", taskID)
	return &api.TaskId{TaskId: taskID}, nil
}
