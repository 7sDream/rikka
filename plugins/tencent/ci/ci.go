package ci

import (
	"github.com/7sDream/rikka/plugins"
	"github.com/tencentyun/image-go-sdk"
)

type TencentCloudImagePlugin struct{}

var (
	l = plugins.SubLogger("[TC-CI]")

	appID      string
	secretID   string
	secretKey  string
	bucketName string
	bucketHost string
	bucketPath string

	// Plugin is the main plugin instance
	Plugin TencentCloudImagePlugin

	cloud *qcloud.PicCloud
)

func buildFullPath(taskID string) string {
	return bucketPath + taskID
}

func (plugin TencentCloudImagePlugin) ExtraHandlers() []plugins.HandlerWithPattern {
	return nil
}
