package plugins

import (
	"flag"
	"net/url"
	"strings"
)

// Common flags for cloud plugins like qiniu and upai
var (
	ArgBucketName = flag.String("bname", "", "Bucket name to store image")
	ArgBucketHost = flag.String("bhost", "", "Bucket host")
	ArgBucketPath = flag.String("bpath", "", "Where the image will be save in bucket")
)

// CheckCommonArgs will check if bname and bhost is set and log their value
func CheckCommonArgs(needName, needHost bool) {
	l.Info("Args bucket name =", *ArgBucketName)
	if needName && *ArgBucketName == "" {
		l.Fatal("No bucket name provided, please add option -bname")
	}

	l.Info("Args bucket host =", *ArgBucketHost)
	if needHost && *ArgBucketHost == "" {
		l.Fatal("No bucket host provided, please add option -bhost")
	}

	l.Info("Args bucket path =", *ArgBucketPath)
}

// GetBucketName get the name of bucket where image will be stored
func GetBucketName() string {
	return *ArgBucketName
}

// GetBucketHost get the host of bucket where image will be stored
func GetBucketHost() string {
	bucketAddr := *ArgBucketHost
	if !strings.HasPrefix(bucketAddr, "http") {
		bucketAddr = "http://" + bucketAddr
	}

	pURL, err := url.Parse(bucketAddr)
	if err != nil {
		l.Fatal("Invalid bucket host", bucketAddr, ":", err)
	}

	bucketAddr = pURL.Scheme + "://" + pURL.Host

	return bucketAddr
}

// GetBucketPath get the save path of bucket which image will be stored to
func GetBucketPath() string {
	bucketPrefix := *ArgBucketPath
	if strings.HasPrefix(bucketPrefix, "/") {
		bucketPrefix = bucketPrefix[1:]
	}
	if len(bucketPrefix) > 0 && !strings.HasSuffix(bucketPrefix, "/") {
		bucketPrefix = bucketPrefix + "/"
	}

	return bucketPrefix
}
