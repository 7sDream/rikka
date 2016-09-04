package fs

import (
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	pathutil "path/filepath"
	"time"

	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

var filesDir = flag.String("dir", "files", "Where files will be save when use fs plugin.")
var tempDir string
var l = plugins.SubLogger("[FS]")

const (
	stateCopying     = "copying"
	stateCopyingCode = 1
	stateCopyingDesc = "Photo is being copied to rikka file system"
)

type fsPlugin struct{}

var FsPlugin fsPlugin = fsPlugin{}

func (fsp fsPlugin) Init() {
	l.Info("Args dir =", *filesDir)
	absFilesDir, err := pathutil.Abs(*filesDir)
	if err == nil {
		l.Info("Abs path of photo file dir:", absFilesDir)
		tempDir = absFilesDir
	} else {
		l.Fatal("A error happened when change photo dir to absolute path:", err)
	}

	if util.CheckExist(absFilesDir) {
		l.Info("Photo file dir already exist")
	} else {
		l.Info("Photo dir not eixst, try to create it")
		err = os.MkdirAll(absFilesDir, 0755)
		if err == nil {
			l.Info("Create dir", absFilesDir, "successfully")
		} else {
			l.Fatal("A error happened when try to create photo dir:", err)
		}
	}

	l.Info("Fs plugin start successfully")
}

func deleteFile(filepath string) {
	if util.CheckExist(filepath) {
		if err := os.Remove(filepath); err != nil {
			l.Fatal("A error happened when try to delete file", filepath, ":", err)
		}
	}
}

func buildCopyingState(taskID string) plugins.State {
	return plugins.State{
		TaskID:      taskID,
		StateCode:   stateCopyingCode,
		State:       stateCopying,
		Description: stateCopyingDesc,
	}
}

func saveFile(uploadFile multipart.File, saveTo *os.File, taskID string) {
	l.Info("Starting task", taskID)

	filepath := saveTo.Name()

	if err := plugins.CreateTask(buildCopyingState(taskID)); err != nil {
		l.Warn("A error happend when add task", taskID, ":", err)
		saveTo.Close()
		uploadFile.Close()
		deleteFile(filepath)
		return
	}

	_, err := io.Copy(saveTo, uploadFile)
	saveTo.Close()
	uploadFile.Close()

	if err == nil {
		// copy file successfully
		l.Info("File copy on task", taskID, "successfully")

		if err := plugins.DeleteTask(taskID); err == nil {
			l.Info("Task", taskID, "finished")
		} else {
			deleteFile(filepath)
			l.Fatal("A error happened when delete task", taskID, ":", err)
		}
	} else {
		// copy file failed
		deleteFile(filepath)
		l.Warn("A error happened when copy file", taskID, ":", err)

		if err := plugins.ChangeTaskState(plugins.BuildErrorState(taskID, err.Error())); err == nil {
			l.Info("Turn task", taskID, "state to error")
		} else {
			l.Fatal("A error happend when change task", taskID, "to error state:", err)
		}
	}
}

func (fsp fsPlugin) SaveRequestHandle(q *plugins.SaveRequest) (taskID string, err error) {
	l.Info("Recieved a file save request")
	now := time.Now().UTC()
	saveTo, err := ioutil.TempFile(tempDir, now.Format("2006-01-02-"))

	if err == nil {
		l.Info("Create file on fs successfully:", saveTo.Name())
	} else {
		l.Warn("Create file on fs error:", err)
		return "", err
	}

	_, taskID = pathutil.Split(saveTo.Name())

	go saveFile(q.File, saveTo, taskID)

	return taskID, nil
}

func (fsp fsPlugin) StateRequestHandle(taskID string) (pState *plugins.State, err error) {
	if pState, err = plugins.GetTaskState(taskID); err == nil {
		return pState, nil
	}
	if util.CheckExist(pathutil.Join(tempDir, taskID)) {
		finishState := plugins.BuildFinishState(taskID)
		return &finishState, nil
	}
	return nil, err
}

func buildURL(r *http.Request, taskID string) string {
	res := url.URL{
		Scheme: "http",
		Host:   r.Host,
		Path:   "files/" + taskID,
	}
	return res.String()
}

func (fsp fsPlugin) URLRequestHandle(q *plugins.URLRequest) (url *plugins.URLJSON, err error) {
	taskID := q.TaskID
	r := q.HTTPRequest
	if util.CheckExist(pathutil.Join(tempDir, taskID)) {
		url := buildURL(r, taskID)
		return &plugins.URLJSON{URL: url}, nil
	}
	return nil, errors.New("File not exist.")
}

func (fsp fsPlugin) ExtraHandlers() (handlers []plugins.HandlerWithPattern) {
	fileFs := http.StripPrefix(
		"/files",
		util.DisableListDir(
			http.FileServer(http.Dir(tempDir)),
		),
	)
	handlers = []plugins.HandlerWithPattern{
		plugins.HandlerWithPattern{
			Pattern: "/files/", Handler: fileFs,
		},
	}

	return handlers
}
