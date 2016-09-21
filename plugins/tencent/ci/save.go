package ci

import (
	"io/ioutil"
	"strings"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/plugins"
	uuid "github.com/satori/go.uuid"
)

const (
	taskIDPlaceholder = "{TaskID}"
)

func uploadToCI(q *plugins.SaveRequest, taskID string) {
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
			plugins.ChangeTaskState(api.BuildErrorState(taskID, errorMsg))
		}
	}()

	err := plugins.ChangeTaskState(buildReadingState(taskID))
	if err != nil {
		l.Fatal("Error happend when change state of task", taskID, "to reading file:", err)
	}
	l.Debug("Change state of task", taskID, "to reading file successfully")

	fileContent, err := ioutil.ReadAll(q.File)
	defer q.File.Close()

	if err != nil {
		l.Error("Error happened when reading uploaded file of task", taskID, ":", err)
		err := plugins.ChangeTaskState(api.BuildErrorState(taskID, err.Error()))
		if err != nil {
			l.Fatal("Error happened when change state of task", taskID, "to error:", err)
		}
		l.Debug("Change state of task", taskID, "to error successfully")
		return
	}
	l.Debug("Read file of task", taskID, "successfully")

	err = plugins.ChangeTaskState(buildUploadingState(taskID))
	if err != nil {
		l.Fatal("Error happend when change state of task", taskID, "to uploading file:", err)
	}
	l.Debug("Change state of task", taskID, "to uploading file successfully")

	info, err := cloud.UploadWithFileid(fileContent, buildFullPath(taskID))
	if err != nil {
		l.Error("Error happened when uploading file of task to ci", taskID, ":", err)
		err := plugins.ChangeTaskState(api.BuildErrorState(taskID, err.Error()))
		if err != nil {
			l.Fatal("Error happened when change state of task", taskID, "to error:", err)
		}
		l.Debug("Change state of task", taskID, "to error successfully")
		return
	}
	l.Debug("Uploading file", taskID, "successfully")

	if err := plugins.DeleteTask(taskID); err != nil {
		l.Error("Error happened when delete task", taskID, ":", err)
	}
	l.Debug("Delete task", taskID, "successfully")

	if bucketHost == "" {
		bucketHost = strings.Replace(info.DownloadUrl, taskID, taskIDPlaceholder, -1)
		l.Debug("Get image url format:", bucketHost)
	}
}

func (cip tcciPlugin) SaveRequestHandle(q *plugins.SaveRequest) (*api.TaskID, error) {
	l.Debug("Recieve a file save request")
	taskID := uuid.NewV4().String() + "." + q.FileExt

	err := plugins.CreateTask(taskID)
	if err != nil {
		l.Fatal("Error happened when create new task!")
	}
	l.Debug("create task", taskID, "successfully, starting background task")

	go uploadToCI(q, taskID)

	l.Debug("Background task started, return task ID:", taskID)
	return &api.TaskID{TaskID: taskID}, nil
}
