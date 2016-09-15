package qiniu

import (
	"net/url"
	"os"
	"strings"

	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
)

// Init is the plugin init function, will be called when plugin be load.
func (qnp qiniuPlugin) Init() {
	l.Info("Start plugin qiniu")

	access = os.Getenv(accessEnvKey)
	secret = os.Getenv(secretEnvKey)

	l.Info("Args access =", util.MaskString(access, 5))
	l.Info("Args secret =", util.MaskString(secret, 5))
	l.Info("Args bucket name =", *plugins.ArgBucketName)
	l.Info("Args bucket host =", *plugins.ArgBucketHost)
	l.Info("Args bucket path =", *plugins.ArgBucketPath)

	if access == "" {
		l.Fatal("No Qiniu access key provided， plesae add it into your env var use the name", accessEnvKey)
	}

	if secret == "" {
		l.Fatal("No Qiniu secret key provided, please add it into your env var use the name", secretEnvKey)
	}

	if *plugins.ArgBucketName == "" {
		l.Fatal("No bucket name provided, please add option -bname")
	}

	if *plugins.ArgBucketName == "" {
		l.Fatal("No bucket host provided, please add option -bhost")
	}

	// host process
	bucketAddr = *plugins.ArgBucketHost
	if !strings.HasPrefix(bucketAddr, "http") {
		bucketAddr = "http://" + bucketAddr
	}

	pURL, err := url.Parse(bucketAddr)
	if err != nil {
		l.Fatal("Invalid bucket host", bucketAddr, ":", err)
	}

	bucketAddr = pURL.Scheme + "://" + pURL.Host

	// name
	bucketName = *plugins.ArgBucketName

	// prefix
	bucketPrefix = *plugins.ArgBucketPath
	if strings.HasPrefix(bucketPrefix, "/") {
		bucketPrefix = bucketPrefix[1:]
	}
	if len(bucketPrefix) > 0 && !strings.HasSuffix(bucketPrefix, "/") {
		bucketPrefix = bucketPrefix + "/"
	}

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
