package cos

import (
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

func (cosp tccosPlugin) Init() {
	l.Info("Start plugin tccos")

	plugins.CheckCommonArgs(true, false)

	appID = util.GetEnvWithCheck("AppID", envAppIDKey, l)
	secretID = util.GetEnvWithCheck("SecretID", envSecretIDKey, l)
	secretKey = util.GetEnvWithCheck("SecretKey", envSecretKeyKey, l)
	bucketName = plugins.GetBucketName()
	bucketPath = plugins.GetBucketPath()

	client = newCosClient()

	l.Info("Tccos plugin start successfully")
}
