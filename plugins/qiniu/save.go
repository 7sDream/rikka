package qiniu

import (
	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/plugins"

	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
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

	// perparing...
	err := plugins.ChangeTaskState(buildPreparingState(taskID))
	if err != nil {
		l.Fatal("Error happend when change state of task", taskID, "to preparing:", err)
	}
	l.Debug("Change state of task", taskID, "to preparing successfully")

	uploader := kodocli.NewUploader(0, nil)
	policy := &kodo.PutPolicy{
		Scope: bucketName,
		//设置Token过期时间
		Expires: 3600,
	}

	// get token
	token, err := client.MakeUptokenWithSafe(policy)
	if err != nil {
		l.Error("Error happened when get Qiniu upload token:", err)
		err = plugins.ChangeTaskState(api.BuildErrorState(taskID, err.Error()))
		if err != nil {
			l.Fatal("Error happened when change state of task", taskID, "to error:", err)
		}
		l.Debug("Change state of task", taskID, "to error sccessfully")
		return
	}

	// uploading
	l.Debug("Upload with arg", "key:", buildPath(taskID), ", filesize:", q.FileSize)
	err = plugins.ChangeTaskState(buildUploadingState(taskID))
	if err != nil {
		l.Fatal("Error happend when change state of task", taskID, "to uploading:", err)
	}
	l.Debug("Change state of task", taskID, "to uploading successfully")

	var ret putRet
	err = uploader.Rput(context.Background(), &ret, token, buildPath(taskID), q.File, q.FileSize, nil)

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

func (qnp qiniuPlugin) SaveRequestHandle(q *plugins.SaveRequest) (*api.TaskID, error) {
	l.Debug("Recieve a file save request")

	taskID := uuid.NewV4().String() + "." + q.FileExt

	err := plugins.CreateTask(taskID)
	if err != nil {
		l.Fatal("Error happened when create new task!")
	}
	l.Debug("create task", taskID, "successfully, starting background task")

	go uploadToQiniu(taskID, q)

	l.Debug("Background task started, return task ID:", taskID)
	return &api.TaskID{TaskID: taskID}, nil
}
