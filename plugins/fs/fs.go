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

var argFilesDir = flag.String("dir", "files", "Where files will be save when use fs plugin.")
var argFsDebugSleep = flag.Int("fsDebugSleep", 0, "Debug: sleep some ms before copy file to fs, used to test javascripta ajax")
var tempDir string
var l = plugins.SubLogger("[FS]")

const (
	stateCopying     = "copying"
	stateCopyingCode = 1
	stateCopyingDesc = "Image is being copied to rikka file system"
)

type fsPlugin struct{}

// FsPlugin is the main plugin instance.
var FsPlugin = fsPlugin{}

// Init is the plugin init function, will be called when plugin be load.
func (fsp fsPlugin) Init() {
	// where to store file
	l.Info("Start plugin fs")
	l.Info("Args dir =", *argFilesDir)
	l.Info("Args fsDebugSleep =", *argFsDebugSleep)

	absFilesDir, err := pathutil.Abs(*argFilesDir)
	if err == nil {
		l.Debug("Abs path of image file dir:", absFilesDir)
		tempDir = absFilesDir
	} else {
		l.Fatal("A error happened when change image dir to absolute path:", err)
	}
	// if target dir not exist, create it
	if util.CheckExist(absFilesDir) {
		l.Debug("Image file dir already exist")
	} else {
		l.Debug("Image file dir not eixst, try to create it")
		err = os.MkdirAll(absFilesDir, 0755)
		if err == nil {
			l.Debug("Create dir", absFilesDir, "successfully")
		} else {
			l.Fatal("A error happened when try to create image dir:", err)
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

	l.Debug("Start copy file of task", taskID)

	// copy file to disk, then close
	_, err := io.Copy(saveTo, uploadFile)
	saveTo.Close()
	uploadFile.Close()

	if err == nil {
		// copy file successfully
		l.Debug("File copy on task", taskID, "finished")

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

		if err := plugins.ChangeTaskState(plugins.BuildErrorState(taskID, err.Error())); err == nil {
			l.Warn("Change task", taskID, "state to error")
		} else {
			// change task state error, exit.
			l.Fatal("A error happend when change task", taskID, "to error state:", err)
		}
	}
}

// SaveRequestHandle Will be called when recieve a file save request.
func (fsp fsPlugin) SaveRequestHandle(q *plugins.SaveRequest) (taskID string, err error) {
	l.Debug("Recieve a file save request")

	// Task ID use time prefix and a number follow it(produce by TempFile)
	now := time.Now().UTC()
	saveTo, err := ioutil.TempFile(tempDir, now.Format("2006-01-02-"))

	if err == nil {
		l.Debug("Create file on fs successfully:", saveTo.Name())
	} else {
		l.Warn("Error happened when try create file:", err)
		return "", err
	}

	_, taskID = pathutil.Split(saveTo.Name())

	// create task
	if plugins.CreateTask(taskID) != nil {
		l.Fatal("Error happened when create new task!")
	}

	// start background copy operate
	go saveFile(q.File, saveTo, taskID)

	l.Debug("Background task started, return task ID:", taskID)
	return taskID, nil
}

// StateRequestHandle Will be called when recieve a get state request.
func (fsp fsPlugin) StateRequestHandle(taskID string) (pState *plugins.State, err error) {

	l.Debug("Recieve a state request of taskID", taskID)

	// taskID exist on task list, just return it
	if pState, err = plugins.GetTaskState(taskID); err == nil {
		l.Debug("State of task", taskID, "found", *pState)
		return pState, nil
	}

	l.Debug("State of task", taskID, "not found, check if file exist")
	// TaskID not exist or error when get it, check if image file already exist
	if util.CheckExist(pathutil.Join(tempDir, taskID)) {
		// file exist as a finished state
		finishState := plugins.BuildFinishState(taskID)
		l.Debug("File of task", taskID, "exist, return finished state", finishState)
		return &finishState, nil
	}

	l.Warn("File of task", taskID, "not exist, get state error:", err)
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

	l.Debug("Recieve an url request of task", taskID)
	l.Debug("Check if file exist of task", taskID)
	// If file exist, return url
	if util.CheckExist(pathutil.Join(tempDir, taskID)) {
		url := buildURL(r, taskID)
		l.Debug("File of task", taskID, "exist, return url", url)
		return &plugins.URLJSON{URL: url}, nil
	}
	l.Error("File of task", taskID, "not exist, return error")
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
