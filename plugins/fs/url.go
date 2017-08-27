package fs

import (
	"errors"
	"net/http"
	"net/url"
	pathUtil "path/filepath"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

// buildURL build complete url from request's Host header and task ID
func buildURL(r *http.Request, taskID string) string {
	res := url.URL{
		Scheme: "http",
		Host:   r.Host,
		//    remove root /
		Path: fileURLPath[1:] + taskID,
	}
	return res.String()
}

// URLRequestHandle will be called when receive a get image url by taskID request
func (fsp fsPlugin) URLRequestHandle(q *plugins.URLRequest) (pURL *api.URL, err error) {
	taskID := q.TaskID
	r := q.HTTPRequest

	l.Debug("Receive an url request of task", taskID)
	l.Debug("Check if file exist of task", taskID)
	// If file exist, return url
	if util.CheckExist(pathUtil.Join(imageDir, taskID)) {
		taskUrl := buildURL(r, taskID)
		l.Debug("File of task", taskID, "exist, return taskUrl", taskUrl)
		return &api.URL{URL: taskUrl}, nil
	}
	l.Error("File of task", taskID, "not exist, return error")
	return nil, errors.New("File not exist.")
}
