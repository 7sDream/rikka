package qiniu

import (
	"net/url"
	"os"
	"strings"
	"unicode/utf8"

	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
)

func maskString(str string, showNum int) string {
	var res string
	var i int
	var c rune
	for i, c = range str {
		if i < showNum {
			res += string(c)
		} else {
			break
		}
	}
	i++
	length := utf8.RuneCountInString(str)
	if i < length {
		res += strings.Repeat("*", length-i)
	}
	return res
}

// Init is the plugin init function, will be called when plugin be load.
func (qnp qiniuPlugin) Init() {
	l.Info("Start plugin qiniu")

	access = os.Getenv(accessEnvKey)
	secret = os.Getenv(secretEnvKey)

	l.Info("Args access =", maskString(access, 5))
	l.Info("Args secret =", maskString(secret, 5))
	l.Info("Args bucket name =", *argBucketName)
	l.Info("Args bucket host =", *argBucketHost)
	l.Info("Args bucket path =", *argBucketPath)

	if access == "" {
		l.Fatal("No Qiniu access key provided， plesae add it into your env var use the name", accessEnvKey)
	}

	if secret == "" {
		l.Fatal("No Qiniu secret key provided, please add it into your env var use the name", secretEnvKey)
	}

	if *argBucketName == "" {
		l.Fatal("No bucket name provided, please add option -bname")
	}

	if *argBucketHost == "" {
		l.Fatal("No bucket host provided, please add option -bhost")
	}

	// host process
	if !strings.HasPrefix(*argBucketHost, "http") {
		*argBucketHost = "http://" + *argBucketHost
	}

	pURL, err := url.Parse(*argBucketHost)
	if err != nil {
		l.Fatal("Invalid bucket host", *argBucketHost, ":", err)
	}

	bucketAddr = pURL.Scheme + "://" + pURL.Host

	// name
	bucketName = *argBucketName

	// disable path arg temporarily, because qiniu sdk has a bug
	if *argBucketPath != "" {
		l.Fatal("The bpath argument is now disabled, plesae don't use it")
	}

	// prefix
	if strings.HasPrefix(*argBucketPath, "/") {
		*argBucketPath = (*argBucketPath)[1:]
	}

	if len(*argBucketPath) > 0 && !strings.HasSuffix(*argBucketPath, "/") {
		*argBucketPath = (*argBucketPath) + "/"
	}
	bucketPrefix = *argBucketPath

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
