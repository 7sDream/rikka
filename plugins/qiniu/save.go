package qiniu

import (
	"context"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/plugins"
	"github.com/qiniu/api.v7/storage"
	"github.com/satori/go.uuid"
)

type putRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

func buildPath(taskID string) string {
	return bucketPrefix + taskID
}

func uploadToQiniu(taskID string, q *plugins.SaveRequest) {
	l.Debug("Getting upload token of task", taskID)

	// preparing...
	err := plugins.ChangeTaskState(buildPreparingState(taskID))
	if err != nil {
		l.Fatal("Error happened when change state of task", taskID, "to preparing:", err)
	}
	l.Debug("Change state of task", taskID, "to preparing successfully")

	policy := storage.PutPolicy{
		Scope: bucketName,
	}
	upToken := policy.UploadToken(mac)

	uploader := storage.NewFormUploader(conf)

	// uploading
	l.Debug("Upload with arg", "key:", buildPath(taskID), ", file size:", q.FileSize)
	err = plugins.ChangeTaskState(buildUploadingState(taskID))
	if err != nil {
		l.Fatal("Error happened when change state of task", taskID, "to uploading:", err)
	}
	l.Debug("Change state of task", taskID, "to uploading successfully")

	var ret putRet
	err = uploader.Put(context.Background(), &ret, upToken, buildPath(taskID), q.File, q.FileSize, nil)

	// uploading error
	if err != nil {
		l.Error("Error happened when upload task", taskID, ":", err)
		err = plugins.ChangeTaskState(api.BuildErrorState(taskID, err.Error()))
		if err != nil {
			l.Fatal("Error happened when change state of task", taskID, "to error:", err)
		}
		l.Debug("Change state of task", taskID, "to error successfully")
	} else {
		// uploading successfully
		l.Info("Upload task", taskID, "to qiniu cloud successfully")
		err = plugins.DeleteTask(taskID)
		if err != nil {
			l.Fatal("Error happened when delete state of task", taskID, ":", err)
		}
		l.Debug("Delete task", taskID, "successfully")
	}
}

func (qnp qiniuPlugin) SaveRequestHandle(q *plugins.SaveRequest) (*api.TaskId, error) {
	l.Debug("Receive a file save request")

	taskID := uuid.NewV4().String() + "." + q.FileExt

	err := plugins.CreateTask(taskID)
	if err != nil {
		l.Fatal("Error happened when create new task!")
	}
	l.Debug("create task", taskID, "successfully, starting background task")

	go uploadToQiniu(taskID, q)

	l.Debug("Background task started, return task ID:", taskID)
	return &api.TaskId{TaskId: taskID}, nil
}
