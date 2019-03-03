package upai

import (
	"github.com/7sDream/rikka/plugins"
	"github.com/upyun/go-sdk/upyun"
)

type upaiPlugin struct{}

const (
	operatorEnvKey = "RIKKA_UPAI_OPERATOR"
	passwordEnvKey = "RIKKA_UPAI_PASSWORD"
)

var (
	l = plugins.SubLogger("[UPai]")

	operator     string
	password     string
	bucketName   string
	bucketAddr   string
	bucketPrefix string

	// Plugin is the main plugin instance
	Plugin = upaiPlugin{}

	client *upyun.UpYun
)

func (qnp upaiPlugin) ExtraHandlers() []plugins.HandlerWithPattern {
	return nil
}
