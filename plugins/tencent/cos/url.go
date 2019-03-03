package cos

import (
	"errors"
	"strings"

	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/plugins"
)

func (plugin tencentCloudObjectStoragePlugin) URLRequestHandle(q *plugins.URLRequest) (*api.URL, error) {
	if bucketHost == "" {
		l.Error("Request URL of task", q.TaskID, "before state become to finish")
		return nil, errors.New("get url before task finish")
	}

	if !strings.Contains(bucketHost, taskIDPlaceholder) {
		l.Fatal("Error! Unable to get image url from bucket host:", bucketHost)
	}

	return &api.URL{
		URL: strings.Replace(bucketHost, taskIDPlaceholder, q.TaskID, -1),
	}, nil
}
