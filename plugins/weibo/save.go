package weibo

import (
	"strconv"
	"sync/atomic"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/plugins"
)

func uploadToWeibo(taskIDInt int64, taskIDStr string, q *plugins.SaveRequest) {
	err := plugins.ChangeTaskState(buildUploadingState(taskIDStr))
	if err != nil {
		l.Fatal("Error happend when change state of task", taskIDStr, "to uploading:", err)
	}

	l.Debug("Uploading to weibo...")
	url, err := upload(client, q)
	if err != nil {
		l.Error("Error happened when upload image to weibo:", err)
		err = plugins.ChangeTaskState(api.BuildErrorState(taskIDStr, err.Error()))
		if err != nil {
			l.Fatal("Error happend when change task", taskIDStr, "to error state:", err)
		}
		return
	}

	urlMap[taskIDInt] = url
	l.Info("Upload task", taskIDStr, "to weibo cloud successfully")

	err = plugins.ChangeTaskState(api.BuildFinishState(taskIDStr))
	if err != nil {
		delete(urlMap, taskIDInt)
		l.Fatal("Error happend when change task", taskIDStr, "to finish state:", err)
	}
}

func (wbp weiboPlugin) SaveRequestHandler(q *plugins.SaveRequest) (*api.TaskID, error) {
	taskIDInt := atomic.AddInt64(&counter, 1)
	taskIDStr := strconv.FormatInt(taskIDInt, 10)

	err := plugins.CreateTask(taskIDStr)
	if err != nil {
		l.Fatal("Error happened when create new task!")
	}

	go uploadToWeibo(taskIDInt, taskIDStr, q)

	return &api.TaskID{TaskID: taskIDStr}, nil
}
