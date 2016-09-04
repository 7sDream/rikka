package fs

import (
	"errors"
	"net/http"
	"net/url"
	pathutil "path/filepath"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

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
func (fsp fsPlugin) URLRequestHandle(q *plugins.URLRequest) (pURL *api.URL, err error) {
	taskID := q.TaskID
	r := q.HTTPRequest

	l.Debug("Recieve an url request of task", taskID)
	l.Debug("Check if file exist of task", taskID)
	// If file exist, return url
	if util.CheckExist(pathutil.Join(imageDir, taskID)) {
		url := buildURL(r, taskID)
		l.Debug("File of task", taskID, "exist, return url", url)
		return &api.URL{URL: url}, nil
	}
	l.Error("File of task", taskID, "not exist, return error")
	return nil, errors.New("File not exist.")
}
