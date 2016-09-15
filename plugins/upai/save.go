package upai

import (
	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/plugins"
	uuid "github.com/satori/go.uuid"
)

func buildPath(taskID string) string {
	return bucketPrefix + taskID
}

func uploadToUPai(taskID string, q *plugins.SaveRequest) {
	// perparing...
	err := plugins.ChangeTaskState(buildUploadingState(taskID))
	if err != nil {
		l.Fatal("Error happend when change state of task", taskID, "to uploading:", err)
	}

	_, err = client.Put(buildPath(taskID), q.File, false, nil)

	if err != nil {
		l.Error("Error happened when upload to upai:", err)
		err = plugins.ChangeTaskState(api.BuildErrorState(taskID, err.Error()))
		if err != nil {
			l.Fatal("Error happened when change state of task", taskID, "to error:", err)
		}
		return
	}
	// uploading successfully
	l.Info("Upload task", taskID, "to upai cloud successfully")
	err = plugins.DeleteTask(taskID)
	if err != nil {
		l.Fatal("Error happened when delete state of task", taskID, ":", err)
	}
}

func (qnp upaiPlugin) SaveRequestHandle(q *plugins.SaveRequest) (*api.TaskID, error) {
	taskID := uuid.NewV4().String() + "." + q.FileExt

	err := plugins.CreateTask(taskID)
	if err != nil {
		l.Fatal("Error happened when create new task!")
	}

	go uploadToUPai(taskID, q)

	return &api.TaskID{TaskID: taskID}, nil
}

func (qnp upaiPlugin) StateRequestHandle(taskID string) (pState *api.State, err error) {
	l.Debug("Recieve a state request of taskID", taskID)

	pState, err = plugins.GetTaskState(taskID)
	if err == nil {
		if pState.StateCode == api.StateErrorCode {
			l.Warn("Get a error state of task", taskID, *pState)
		} else {
			l.Debug("Get a normal state of task", taskID, *pState)
		}
		return pState, nil
	}

	l.Debug("State of task", taskID, "not found, just return a finish state")
	return api.BuildFinishState(taskID), nil
}
