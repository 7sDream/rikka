package fs

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	pathutil "path/filepath"
	"time"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

func deleteFile(filepath string) {
	if util.CheckExist(filepath) {
		if err := os.Remove(filepath); err != nil {
			l.Fatal("A error happened when try to delete file", filepath, ":", err)
		}
	}
}

// background operate, save file to disk
func saveFile(uploadFile multipart.File, saveTo *os.File, taskID string) {
	filepath := saveTo.Name()

	// If error happend when change task state, delete file
	if err := plugins.ChangeTaskState(buildCopyingState(taskID)); err != nil {
		l.Error("Error happend when change state of task", taskID, "to copying:", err)
		saveTo.Close()
		uploadFile.Close()
		deleteFile(filepath)
		return
	}

	l.Debug("Change task", taskID, "state to copy successfully")

	// sleep for debug javascripta
	if *argFsDebugSleep > 0 {
		l.Debug("Sleep", *argFsDebugSleep, "ms for debug")
		time.Sleep(time.Duration(*argFsDebugSleep) * time.Millisecond)
	}

	l.Info("Start copy file of task", taskID)

	// copy file to disk, then close
	_, err := io.Copy(saveTo, uploadFile)
	saveTo.Close()
	uploadFile.Close()

	if err == nil {
		// copy file successfully
		l.Info("File copy on task", taskID, "finished")

		if err := plugins.DeleteTask(taskID); err == nil {
			l.Debug("Task", taskID, "finished, deleted it from task list")
		} else {
			// delete task failed, delete file and exit
			deleteFile(filepath)
			l.Fatal("A error happened when delete task", taskID, ":", err)
		}
	} else {
		// copy file failed, delete file and turn state to error
		l.Warn("Error happened when copy file of task", taskID, ":", err)
		deleteFile(filepath)

		if err := plugins.ChangeTaskState(api.BuildErrorState(taskID, err.Error())); err == nil {
			l.Warn("Change task", taskID, "state to error")
		} else {
			// change task state error, exit.
			l.Fatal("A error happend when change task", taskID, "to error state:", err)
		}
	}
}

// SaveRequestHandle Will be called when recieve a file save request.
func (fsp fsPlugin) SaveRequestHandle(q *plugins.SaveRequest) (*api.TaskID, error) {
	l.Debug("Recieve a file save request")

	// Task ID use time prefix and a number follow it(produce by TempFile)
	now := time.Now().UTC()
	saveTo, err := ioutil.TempFile(imageDir, now.Format("2006-01-02-"))

	if err == nil {
		l.Debug("Create file on fs successfully:", saveTo.Name())
	} else {
		l.Warn("Error happened when try create file:", err)
		return nil, err
	}

	_, taskID := pathutil.Split(saveTo.Name())

	// create task
	if plugins.CreateTask(taskID) != nil {
		l.Fatal("Error happened when create new task!")
	}

	// start background copy operate
	go saveFile(q.File, saveTo, taskID)

	l.Debug("Background task started, return task ID:", taskID)
	return &api.TaskID{TaskID: taskID}, nil
}
