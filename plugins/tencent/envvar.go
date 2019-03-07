package tencent

import (
	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/common/util"
)

const (
	envAppIDKey     = "RIKKA_TENCENT_APPID"
	envSecretIDKey  = "RIKKA_TENCENT_SECRETID"
	envSecretKeyKey = "RIKKA_TENCENT_SECRETKEY"
	envRegionKey    = "RIKKA_TENCENT_REGION"
)

// GetAppIDWithCheck will get APPID of  Tencent Cloud.
// If it is empty, will raise a Fatal.
func GetAppIDWithCheck(l *logger.Logger) string {
	return util.GetEnvWithCheck("AppID", envAppIDKey, l)
}

// GetSecretIDWithCheck will get SecretID of  Tencent Cloud.
// If it is empty, will raise a Fatal.
func GetSecretIDWithCheck(l *logger.Logger) string {
	return util.GetEnvWithCheck("SecretID", envSecretIDKey, l)
}

// GetSecretKeyWithCheck will get SecretKey of  Tencent Cloud.
// If it is empty, will raise a Fatal.
func GetSecretKeyWithCheck(l *logger.Logger) string {
	return util.GetEnvWithCheck("SecretKey", envSecretKeyKey, l)
}

// GetRegionWithCheck will get Region of Tencent Cloud.
// If it is empty, will raise a Fatal
func GetRegionWithCheck(l *logger.Logger) string {
	return util.GetEnvWithCheck("Region", envRegionKey, l)
}
