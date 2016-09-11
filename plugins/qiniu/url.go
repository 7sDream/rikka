package qiniu

import (
	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/plugins"
)

func buildURL(taskID string) string {
	return bucketAddr + "/" + buildPath(taskID)
}

// URLRequestHandle will be called when recieve a get image url by taskID request
func (qnp qiniuPlugin) URLRequestHandle(q *plugins.URLRequest) (pURL *api.URL, err error) {
	l.Debug("Recieve an url request of task", q.TaskID)
	return &api.URL{URL: buildURL(q.TaskID)}, nil
}
