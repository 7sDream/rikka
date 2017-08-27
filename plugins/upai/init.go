package upai

import (
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
	"github.com/upyun/go-sdk/upyun"
)

func (qnp upaiPlugin) Init() {
	l.Info("Start plugin upai")

	plugins.CheckCommonArgs(true, true)

	operator = util.GetEnvWithCheck("Operator", operatorEnvKey, l)
	password = util.GetEnvWithCheck("Password", passwordEnvKey, l)
	bucketName = plugins.GetBucketName()
	bucketAddr = plugins.GetBucketHost()
	bucketPrefix = plugins.GetBucketPath()

	client = upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   bucketName,
		Operator: operator,
		Password: password,
	})

	l.Info("UPai plugin start successfully")
}
