package cos

import "github.com/7sDream/rikka/plugins"

type tencentCloudObjectStoragePlugin struct{}

var (
	l = plugins.SubLogger("[TC-COS]")

	appID      string
	secretID   string
	secretKey  string
	bucketName string
	bucketPath string
	bucketHost string

	client *cosClient

	// Plugin is the main plugin instance
	Plugin tencentCloudObjectStoragePlugin
)

func (plugin tencentCloudObjectStoragePlugin) ExtraHandlers() []plugins.HandlerWithPattern {
	return nil
}
