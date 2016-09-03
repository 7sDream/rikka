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

const stateCopyingCode = 1
const stateCopying = "copying"
const stateCopyingDesc = "Photo is being copied to file system"

const stateErrorCode = 999
const stateError = "error"

type fsPlugin struct{}

var FsPlugin fsPlugin = fsPlugin{}

func (fsp fsPlugin) Init() {
	l.Info("Get Photo dir from arguments:", *filesDir)
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

func saveFile(uploadFile multipart.File, saveTo *os.File, fileID string) {
	defer uploadFile.Close()
	defer saveTo.Close()

	l.Info("Starting task", fileID)

	if err := plugins.CreateTask(fileID, stateCopyingCode, stateCopying, stateCopyingDesc); err != nil {
		l.Error("A error happend when add task", fileID, ":", err)
		if util.CheckExist(saveTo.Name()) {
			if err = os.Remove(saveTo.Name()); err != nil {
				l.Error("A error happened when try to delete tempfile", saveTo.Name(), ":", err)
			}
		}
		return
	}

	if _, err := io.Copy(saveTo, uploadFile); err == nil {
		l.Info("File copy on task", fileID, "successfully")
		if err := plugins.FinishTask(fileID); err == nil {
			l.Info("Task", fileID, "finished")
		} else {
			l.Error("A error happened when finish task", fileID, ":", err)
		}
	} else {
		l.Error("A error happened when copy file", fileID, ":", err)
		if err := plugins.ChangeTaskState(fileID, stateErrorCode, stateError, err.Error()); err != nil {
			l.Info("A error happend when change task", fileID, "state:", err)
		}
	}
}

func (fsp fsPlugin) SaveRequestHandle(q *plugins.SaveRequest) (response *plugins.SaveResponse, err error) {
	l.Info("Recieved a file save request")
	now := time.Now().UTC()
	saveTo, err := ioutil.TempFile(tempDir, now.Format("2006-01-02-"))
	if err == nil {
		l.Info("Create file on fs successfully:", saveTo.Name())
	} else {
		l.Error("Create file on fs error:", err)
		return nil, err
	}

	_, name := pathutil.Split(saveTo.Name())

	go saveFile(q.File, saveTo, name)

	return &plugins.SaveResponse{TaskID: name}, nil
}

func (fsp fsPlugin) StateRequestHandle(q *plugins.StateRequest) (response *plugins.StateResponse, err error) {
	taskID := q.TaskID
	if pState, err := plugins.GetTaskState(taskID); err == nil {
		return &plugins.StateResponse{State: *pState}, nil
	}
	return nil, nil
}

func buildURL(r *http.Request, taskID string) string {
	res := url.URL{
		Scheme: "http",
		Host:   r.Host,
		Path:   "files/" + taskID,
	}
	return res.String()
}

func (fsp fsPlugin) GetSrcURL(q *plugins.SrcURLRequest) (url string, err error) {
	taskID := q.TaskID
	r := q.HttpRequest
	if util.CheckExist(taskID) {
		url := buildURL(r, taskID)
		return url, nil
	} else {
		return "", errors.New("File not exist.")
	}
}

func (fsp fsPlugin) ExtraHandlers() (handlers []plugins.HandlerWithPattern) {
	fileFs := http.StripPrefix(
		"/files",
		util.DisableListDir(
			http.FileServer(http.Dir("files")),
		),
	)
	handlers = []plugins.HandlerWithPattern{
		plugins.HandlerWithPattern{
			Pattern: "/files/", Handler: fileFs,
		},
	}

	return handlers
}
