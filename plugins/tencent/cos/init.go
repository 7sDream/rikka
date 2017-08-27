package cos

import (
	"github.com/7sDream/rikka/plugins"
	"github.com/7sDream/rikka/plugins/tencent"
)

func (plugin tccosPlugin) Init() {
	l.Info("Start plugin tccos")

	plugins.CheckCommonArgs(true, false)

	appID = tencent.GetAppIDWithCheck(l)
	secretID = tencent.GetSecretIDWithCheck(l)
	secretKey = tencent.GetSecretKeyWithCheck(l)
	bucketName = plugins.GetBucketName()
	bucketPath = plugins.GetBucketPath()

	client = newCosClient()

	l.Info("Tccos plugin start successfully")
}
