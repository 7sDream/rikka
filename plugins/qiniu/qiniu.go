package qiniu

import (
	"flag"
	"github.com/7sDream/rikka/plugins"
	"qiniupkg.com/api.v7/kodo"
)

// plugin type
type qiniuPlugin struct{}

const (
	accessEnvKey = "RIKKA_QINIU_ACCESS"
	secretEnvKey = "RIKKA_QINIU_SECRET"
)

var (
	l = plugins.SubLogger("[Qiniu]")

	argBucketName = flag.String("bname", "", "Qiniu bucket name to store image")
	argBucketHost = flag.String("bhost", "", "Qiniu bucket host")
	argBucketPath = flag.String("bpath", "", "Where the image will be save in Qiniu bucket")

	access       string
	secret       string
	bucketName   string
	bucketAddr   string
	bucketPrefix string
	client       *kodo.Client

	// QiniuPlugin is the main plugin instance
	QiniuPlugin = qiniuPlugin{}
)

func (qnp qiniuPlugin) ExtraHandlers() []plugins.HandlerWithPattern {
	return nil
}
