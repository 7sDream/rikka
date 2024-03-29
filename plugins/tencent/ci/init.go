package ci

import (
	"strconv"

	"github.com/7sDream/rikka/plugins"
	"github.com/7sDream/rikka/plugins/tencent"
	qcloud "github.com/jeremywangjun/image-go-sdk"
)

func (plugin TencentCloudImagePlugin) Init() {
	l.Info("Start plugin tencent cloud image")

	plugins.CheckCommonArgs(true, false)

	appID = tencent.GetAppIDWithCheck(l)
	secretID = tencent.GetSecretIDWithCheck(l)
	secretKey = tencent.GetSecretKeyWithCheck(l)
	bucketName = plugins.GetBucketName()
	bucketPath = plugins.GetBucketPath()

	appIDUint, err := strconv.Atoi(appID)
	if err != nil {
		l.Fatal("Error happened when parse APPID to int:", err)
	}
	l.Debug("Parse APPID to int successfully")

	cloud = &qcloud.PicCloud{
		Appid:     uint(appIDUint),
		SecretId:  secretID,
		SecretKey: secretKey,
		Bucket:    bucketName,
	}

	l.Info("Tencent cloud image plugin start successfully")
}
