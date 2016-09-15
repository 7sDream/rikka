package qiniu

import (
	"os"

	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
)

// Init is the plugin init function, will be called when plugin be load.
func (qnp qiniuPlugin) Init() {
	l.Info("Start plugin qiniu")

	plugins.CheckCommonArgs()

	access = os.Getenv(accessEnvKey)
	secret = os.Getenv(secretEnvKey)

	l.Info("Args access =", util.MaskString(access, 5))
	l.Info("Args secret =", util.MaskString(secret, 5))

	if access == "" {
		l.Fatal("No Qiniu access key provided， plesae add it into your env var use the name", accessEnvKey)
	}

	if secret == "" {
		l.Fatal("No Qiniu secret key provided, please add it into your env var use the name", secretEnvKey)
	}

	// name
	bucketName = plugins.GetBucketName()

	// host process
	bucketAddr = plugins.GetBucketHost()

	// path
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
