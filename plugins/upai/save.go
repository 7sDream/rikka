package upai

import (
	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/plugins"
	"github.com/satori/go.uuid"
	"github.com/upyun/go-sdk/v3/upyun"
)

func buildPath(taskID string) string {
	return bucketPrefix + taskID
}

func uploadToUPai(taskID string, q *plugins.SaveRequest) {
	// preparing...
	err := plugins.ChangeTaskState(buildUploadingState(taskID))
	if err != nil {
		l.Fatal("Error happen when change state of task", taskID, "to uploading:", err)
	}
	l.Debug("Change state of task", taskID, "to uploading successfully")

	l.Debug("Uploading to UPai cloud...")

	err = client.Put(&upyun.PutObjectConfig{
		Path:            buildPath(taskID),
		Reader:          q.File,
		UseResumeUpload: false,
	})

	if err != nil {
		l.Error("Error happened when upload to upai:", err)
		err = plugins.ChangeTaskState(api.BuildErrorState(taskID, err.Error()))
		if err != nil {
			l.Fatal("Error happened when change state of task", taskID, "to error:", err)
		}
		l.Debug("Change state of task", taskID, "to error successfully")
		return
	}
	// uploading successfully
	l.Info("Upload task", taskID, "to upai cloud successfully")

	err = plugins.DeleteTask(taskID)
	if err != nil {
		l.Fatal("Error happened when delete state of task", taskID, ":", err)
	}
	l.Debug("Delete task", taskID, "successfully")
}

func (qnp upaiPlugin) SaveRequestHandle(q *plugins.SaveRequest) (*api.TaskId, error) {
	l.Debug("Receive a file save request")

	taskID := uuid.NewV4().String() + "." + q.FileExt

	err := plugins.CreateTask(taskID)
	if err != nil {
		l.Fatal("Error happened when create new task!")
	}
	l.Debug("create task", taskID, "successfully, starting background task")

	go uploadToUPai(taskID, q)

	l.Debug("Background task started, return task ID:", taskID)
	return &api.TaskId{TaskId: taskID}, nil
}
