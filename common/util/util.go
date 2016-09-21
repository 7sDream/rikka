package util

import (
	"os"

	"github.com/7sDream/rikka/common/logger"
)

var (
	l = logger.NewLogger("[Util]")
)

// GetEnvWithCheck will get a env var, print it, and return it.
// If the var is empty, will raise a Fatal.
func GetEnvWithCheck(name, key string, log *logger.Logger) string {
	if log == nil {
		log = l
	}
	value := os.Getenv(key)
	if value == "" {
		log.Fatal("No", name, "provided, please add it to your env var use the name", key)
	}
	log.Info("Args", name, "=", MaskString(value, 5))
	return value
}
