package qiniu

import (
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
	"github.com/qiniu/api.v7/auth/qbox"
)

// Init is the plugin init function, will be called when plugin be load.
func (qnp qiniuPlugin) Init() {
	l.Info("Start plugin qiniu")

	plugins.CheckCommonArgs(true, true)

	access = util.GetEnvWithCheck("Access", accessEnvKey, l)
	secret = util.GetEnvWithCheck("Secret", secretEnvKey, l)
	bucketName = plugins.GetBucketName()
	bucketAddr = plugins.GetBucketHost()
	bucketPrefix = plugins.GetBucketPath()

	mac = qbox.NewMac(access, secret)

	l.Info("Qiniu plugin start successfully")
}
