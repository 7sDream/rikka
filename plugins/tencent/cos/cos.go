package cos

import "github.com/7sDream/rikka/plugins"

type tccosPlugin struct{}

const (
	envAppIDKey     = "RIKKA_TCCOS_APPID"
	envSecretIDKey  = "RIKKA_TCCOS_SECRETID"
	envSecretKeyKey = "RIKKA_TCCOS_SECRETKEY"
)

var (
	l = plugins.SubLogger("[TCcos]")

	appID      string
	secretID   string
	secretKey  string
	bucketName string
	bucketPath string
	bucketHost string

	client *cosClient

	// TCcosPlugin is the main plugin instance
	TCcosPlugin tccosPlugin
)

func (cosp tccosPlugin) ExtraHandlers() []plugins.HandlerWithPattern {
	return nil
}
