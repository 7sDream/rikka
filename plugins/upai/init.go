package upai

import (
	"os"

	"github.com/7sDream/rikka/plugins"
	"github.com/upyun/go-sdk/upyun"
)

func (qnp upaiPlugin) Init() {
	l.Info("Start plugin upai")

	plugins.CheckCommonArgs()

	operator = os.Getenv(operatorEnvKey)
	password = os.Getenv(passwordEnvKey)

	if operator == "" {
		l.Fatal("No UPai operator name provided， plesae add it into your env var use the name", operatorEnvKey)
	}

	if password == "" {
		l.Fatal("No UPai password provided, please add it into your env var use the name", passwordEnvKey)
	}

	// name
	bucketName = plugins.GetBucketName()

	// host
	bucketAddr = plugins.GetBucketHost()

	// path
	bucketPrefix = plugins.GetBucketPath()

	client = upyun.NewUpYun(bucketName, operator, password)

	l.Info("UPai plugin start successfully")
}
