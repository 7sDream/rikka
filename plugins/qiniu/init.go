package qiniu

import (
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
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

	// set qiniu conf
	conf.ACCESS_KEY = access
	conf.SECRET_KEY = secret
	config := &kodo.Config{
		AccessKey: access,
		SecretKey: secret,
	}
	client = kodo.New(0, config)

	l.Info("Qiniu plugin start successfully")
}
