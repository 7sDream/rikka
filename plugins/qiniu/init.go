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

	if access == "" {
		l.Fatal("No Qiniu access key provided")
	}

	if secret == "" {
		l.Fatal("No Qiniu secret key provided")
	}

	if *argBucketName == "" {
		l.Fatal("No bucket name provided")
	}

	if *argBucketHost == "" {
		l.Fatal("No bucket host provided")
	}

	if !strings.HasPrefix(*argBucketHost, "http") {
		*argBucketHost = "http://" + *argBucketHost
	}

	pURL, err := url.Parse(*argBucketHost)
	if err != nil {
		l.Fatal("Invalid bucket host", *argBucketHost, ":", err)
	}

	bucketName = *argBucketName
	bucketAddr = pURL.Scheme + "://" + pURL.Host

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
