package cos

import (
	"github.com/7sDream/rikka/plugins"
	"github.com/7sDream/rikka/plugins/tencent"
)

func (plugin tencentCloudObjectStoragePlugin) Init() {
	l.Info("Start plugin tencent cloud object storage")

	plugins.CheckCommonArgs(true, false)

	appID = tencent.GetAppIDWithCheck(l)
	secretID = tencent.GetSecretIDWithCheck(l)
	secretKey = tencent.GetSecretKeyWithCheck(l)
	bucketName = plugins.GetBucketName()
	bucketPath = plugins.GetBucketPath()

	client = newCosClient()

	l.Info("Tencent cloud object storage plugin start successfully")
}
