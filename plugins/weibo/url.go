package weibo

import (
	"strconv"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/plugins"
)

func (wbp weiboPlugin) URLRequestHandle(q *plugins.URLRequest) (pURL *api.URL, err error) {
	taskIDInt, err := strconv.ParseInt(q.TaskID, 10, 64)
	if err != nil {
		l.Fatal("Error happened when parser taskID", q.TaskID, "to int64:", err)
	}
	url, ok := urlMap[taskIDInt]
	if !ok {
		l.Fatal("Can't get url of a finshed task", q.TaskID)
	}
	return &api.URL{URL: url}, nil
}
