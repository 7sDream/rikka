package fs

import (
	"io"
	"mime/multipart"
	"os"
	pathutil "path/filepath"
	"time"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"

	"github.com/satori/go.uuid"
)

func deleteFile(filepath string) {
	if util.CheckExist(filepath) {
		if err := os.Remove(filepath); err != nil {
			l.Fatal("A error happened when try to delete file", filepath, ":", err)
		}
	}
}

func createFile(uploadFile multipart.File, filepath string, taskID string) (*os.File, error) {
	// If error happend when change task state, close file
	if err := plugins.ChangeTaskState(buildCreatingState(taskID)); err != nil {
		uploadFile.Close()
		l.Fatal("Error happend when change state of task", taskID, "to copying:", err)
	}
	l.Debug("Change state of task", taskID, "to creating state successfully")

	saveTo, err := os.Create(filepath)

	if err != nil {
		// create file failed, close file and change state
		l.Error("Error happened when create file of task", taskID, ":", err)
		uploadFile.Close()

		if err := plugins.ChangeTaskState(api.BuildErrorState(taskID, err.Error())); err != nil {
			// change task state error, exit.
			l.Fatal("A error happend when change task", taskID, "to error state:", err)
		} else {
			l.Warn("Change task", taskID, "state to error successfully")
		}
		return nil, err
	}

	l.Debug("Create file on fs successfully:", saveTo.Name())

	return saveTo, nil
}

func fileCopy(uploadFile multipart.File, saveTo *os.File, filepath string, taskID string) {
	// If error happend when change task state, delete file and close
	if err := plugins.ChangeTaskState(buildCopyingState(taskID)); err != nil {
		uploadFile.Close()
		saveTo.Close()
		deleteFile(filepath)
		l.Fatal("Error happend when change state of task", taskID, "to copying:", err)
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

	if err != nil {
		// copy file failed, delete file and turn state to error
		l.Warn("Error happened when copy file of task", taskID, ":", err)
		deleteFile(filepath)

		var err2 error
		if err2 = plugins.ChangeTaskState(api.BuildErrorState(taskID, err.Error())); err2 != nil {
			l.Fatal("Error happend when change task", taskID, "to error state:", err2)
		} else {
			l.Warn("Change task", taskID, "state to error successfully")
		}
	}

	// copy file successfully
	l.Info("File copy of task", taskID, "finished")

	// delete successful task, non-exist task means successful
	if err := plugins.DeleteTask(taskID); err == nil {
		l.Debug("Task", taskID, "finished, deleted it from task list")
	} else {
		// delete task failed, delete file and exit
		deleteFile(filepath)
		l.Fatal("A error happened when delete task", taskID, ":", err)
	}
}

// background operate, save file to disk
func saveFile(uploadFile multipart.File, filename string) {
	filepath := pathutil.Join(imageDir, filename)

	saveTo, err := createFile(uploadFile, filepath, filename)
	if err != nil {
		return
	}

	fileCopy(uploadFile, saveTo, filepath, filename)
}

// SaveRequestHandle Will be called when recieve a file save request.
func (fsp fsPlugin) SaveRequestHandle(q *plugins.SaveRequest) (*api.TaskID, error) {
	l.Debug("Recieve a file save request")

	taskID := uuid.NewV4().String() + "." + q.FileExt

	// create task
	if plugins.CreateTask(taskID) != nil {
		l.Fatal("Error happened when create new task!")
	}
	l.Debug("create task", taskID, "successfully, starting background task")

	// start background copy operate
	go saveFile(q.File, taskID)

	l.Debug("Background task started, return task ID:", taskID)
	return &api.TaskID{TaskID: taskID}, nil
}
