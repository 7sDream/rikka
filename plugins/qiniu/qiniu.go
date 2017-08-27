package qiniu

import (
	"github.com/7sDream/rikka/plugins"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

// plugin type
type qiniuPlugin struct{}

const (
	accessEnvKey = "RIKKA_QINIU_ACCESS"
	secretEnvKey = "RIKKA_QINIU_SECRET"
)

var (
	l = plugins.SubLogger("[Qiniu]")

	access       string
	secret       string
	bucketName   string
	bucketAddr   string
	bucketPrefix string

	conf = &storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseHTTPS:      true,
		UseCdnDomains: true,
	}
	mac *qbox.Mac

	// QiniuPlugin is the main plugin instance
	QiniuPlugin = qiniuPlugin{}
)

func (qnp qiniuPlugin) ExtraHandlers() []plugins.HandlerWithPattern {
	return nil
}
