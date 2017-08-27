package weibo

import (
	"strconv"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/plugins"
)

func (wbp weiboPlugin) URLRequestHandle(q *plugins.URLRequest) (pURL *api.URL, err error) {
	l.Debug("Receive an url request of task", q.TaskID)

	taskIDInt, err := strconv.ParseInt(q.TaskID, 10, 64)
	if err != nil {
		l.Fatal("Error happened when parser taskID", q.TaskID, "to int64:", err)
	}
	l.Debug("Parse int from task ID", q.TaskID, "successfully")

	imageID, ok := imageIDMap[taskIDInt]
	if !ok {
		l.Fatal("Can't get url of a finished task", q.TaskID)
	}
	imageURL := buildURL(imageID)
	l.Debug("Get image ID", imageID, "successfully, return url", imageURL)

	return &api.URL{URL: imageURL}, nil
}
