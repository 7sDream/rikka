package weibo

import (
	"strconv"
	"sync/atomic"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/plugins"
)

func uploadToWeibo(taskIDInt int64, taskIDStr string, q *plugins.SaveRequest) {
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
			plugins.ChangeTaskState(api.BuildErrorState(taskIDStr, errorMsg))
		}
	}()

	err := plugins.ChangeTaskState(buildUploadingState(taskIDStr))
	if err != nil {
		l.Fatal("Error happened when change state of task", taskIDStr, "to uploading:", err)
	}
	l.Debug("Change state of task", taskIDStr, "to uploading successfully")

	l.Debug("Uploading to weibo...")
	url, err := upload(q)
	if err != nil {
		l.Error("Error happened when upload image to weibo:", err)
		err = plugins.ChangeTaskState(api.BuildErrorState(taskIDStr, err.Error()))
		if err != nil {
			l.Fatal("Error happened when change task", taskIDStr, "to error state:", err)
		}
		l.Debug("Change state of task", taskIDStr, "to error successfully")
		return
	}

	imageIDMap[taskIDInt] = url
	l.Info("Upload task", taskIDStr, "to weibo cloud successfully")

	err = plugins.DeleteTask(taskIDStr)
	if err != nil {
		delete(imageIDMap, taskIDInt)
		l.Fatal("Error happened when change task", taskIDStr, "to finish state:", err)
	}
	l.Debug("Delete task", taskIDStr, "successfully")
}

func (wbp weiboPlugin) SaveRequestHandle(q *plugins.SaveRequest) (*api.TaskId, error) {
	l.Debug("Receive a file save request")

	taskIDInt := atomic.AddInt64(&counter, 1)
	taskIDStr := strconv.FormatInt(taskIDInt, 10)

	err := plugins.CreateTask(taskIDStr)
	if err != nil {
		l.Fatal("Error happened when create new task!")
	}
	l.Debug("create task", taskIDStr, "successfully, starting background task")

	go uploadToWeibo(taskIDInt, taskIDStr, q)

	l.Debug("Background task started, return task ID:", taskIDStr)
	return &api.TaskId{TaskId: taskIDStr}, nil
}
