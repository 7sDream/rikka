package ci

import (
	"github.com/7sDream/rikka/plugins"
	qcloud "github.com/tencentyun/image-go-sdk"
)

type tcciPlugin struct{}

var (
	l = plugins.SubLogger("[TC-CI]")

	appID      string
	secretID   string
	secretKey  string
	bucketName string
	bucketHost string
	bucketPath string

	// TCciPlugin is the main plugin instance
	TCciPlugin tcciPlugin

	cloud *qcloud.PicCloud
)

func buildFullPath(taskID string) string {
	return bucketPath + taskID
}

func (cip tcciPlugin) ExtraHandlers() []plugins.HandlerWithPattern {
	return nil
}
