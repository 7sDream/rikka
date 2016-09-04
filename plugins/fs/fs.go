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

// FsPlugin is the main plugin instance.
var FsPlugin = fsPlugin{}

// Init is the plugin init function, will be called when plugin be load.
func (fsp fsPlugin) Init() {
	// where to store file
	l.Info("Args dir =", *filesDir)
	absFilesDir, err := pathutil.Abs(*filesDir)
	if err == nil {
		l.Info("Abs path of photo file dir:", absFilesDir)
		tempDir = absFilesDir
	} else {
		l.Fatal("A error happened when change photo dir to absolute path:", err)
	}
	// if target dir not exist, create it
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

// A shortcut funtion to build state we need.
func buildCopyingState(taskID string) plugins.State {
	return plugins.State{
		TaskID:      taskID,
		StateCode:   stateCopyingCode,
		State:       stateCopying,
		Description: stateCopyingDesc,
	}
}

// background operate, save file to disk
func saveFile(uploadFile multipart.File, saveTo *os.File, taskID string) {
	l.Info("Starting task", taskID)

	filepath := saveTo.Name()

	// If error happend when create task, delete file
	if err := plugins.CreateTask(buildCopyingState(taskID)); err != nil {
		l.Warn("A error happend when add task", taskID, ":", err)
		saveTo.Close()
		uploadFile.Close()
		deleteFile(filepath)
		return
	}

	// copy file to disk, then close
	_, err := io.Copy(saveTo, uploadFile)
	saveTo.Close()
	uploadFile.Close()

	if err == nil {
		// copy file successfully
		l.Info("File copy on task", taskID, "successfully")

		if err := plugins.DeleteTask(taskID); err == nil {
			l.Info("Task", taskID, "finished")
		} else {
			// delete task failed, delete file and exit
			deleteFile(filepath)
			l.Fatal("A error happened when delete task", taskID, ":", err)
		}
	} else {
		// copy file failed, delete file and turn state to error
		deleteFile(filepath)
		l.Warn("A error happened when copy file", taskID, ":", err)

		if err := plugins.ChangeTaskState(plugins.BuildErrorState(taskID, err.Error())); err == nil {
			l.Info("Turn task", taskID, "state to error")
		} else {
			// change task state error, exit.
			l.Fatal("A error happend when change task", taskID, "to error state:", err)
		}
	}
}

// SaveRequestHandle Will be called when recieve a file save request.
func (fsp fsPlugin) SaveRequestHandle(q *plugins.SaveRequest) (taskID string, err error) {
	l.Info("Recieved a file save request")

	// Task ID use time prefix and a number follow it(produce by TempFile)
	now := time.Now().UTC()
	saveTo, err := ioutil.TempFile(tempDir, now.Format("2006-01-02-"))

	if err == nil {
		l.Info("Create file on fs successfully:", saveTo.Name())
	} else {
		l.Error("Create file on fs error:", err)
		return "", err
	}

	_, taskID = pathutil.Split(saveTo.Name())

	// start background copy operate
	go saveFile(q.File, saveTo, taskID)

	return taskID, nil
}

// StateRequestHandle Will be called when recieve a get state request.
func (fsp fsPlugin) StateRequestHandle(taskID string) (pState *plugins.State, err error) {
	// taskID exist on task list, just return it
	if pState, err = plugins.GetTaskState(taskID); err == nil {
		return pState, nil
	}
	// TaskID not exist or error when get it, check if image file already exist
	if util.CheckExist(pathutil.Join(tempDir, taskID)) {
		// file exist as a finished state
		finishState := plugins.BuildFinishState(taskID)
		return &finishState, nil
	}
	// get state error
	return nil, err
}

// buildURL build complete url from request's Host header and task ID
func buildURL(r *http.Request, taskID string) string {
	res := url.URL{
		Scheme: "http",
		Host:   r.Host,
		Path:   "files/" + taskID,
	}
	return res.String()
}

// URLRequestHandle will be called when recieve a get image url by taskID request
func (fsp fsPlugin) URLRequestHandle(q *plugins.URLRequest) (url *plugins.URLJSON, err error) {
	taskID := q.TaskID
	r := q.HTTPRequest
	// If file exist, return url
	if util.CheckExist(pathutil.Join(tempDir, taskID)) {
		url := buildURL(r, taskID)
		return &plugins.URLJSON{URL: url}, nil
	}
	return nil, errors.New("File not exist.")
}

// ExtraHandlers return value will be add to http handle list.
// In fs plugin, we start a static file server to serve image file we accped in /files/taskID path.
func (fsp fsPlugin) ExtraHandlers() (handlers []plugins.HandlerWithPattern) {
	// get a base file server
	fileServer := http.StripPrefix(
		"/files",
		// Disable list dir
		util.DisableListDir(
			http.FileServer(http.Dir(tempDir)),
		),
	)
	// only accped GET request
	requestFilterFileServer := util.RequestFilter(
		"", "GET", l,
		func(w http.ResponseWriter, q *http.Request) {
			fileServer.ServeHTTP(w, q)
		},
	)

	handlers = []plugins.HandlerWithPattern{
		plugins.HandlerWithPattern{
			Pattern: "/files/", Handler: requestFilterFileServer,
		},
	}

	return handlers
}
