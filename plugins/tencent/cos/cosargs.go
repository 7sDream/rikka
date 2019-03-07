package cos

import (
	"flag"
	"github.com/7sDream/rikka/common/logger"
)

// flags for Tencent cloud cos sdk
var ArgCosVersion = flag.String("bVersion", "", "Tencent cos sdk version")

// GetVersion get the version to know the version of Tencent cos version
func GetVersionWitchCheck(l *logger.Logger) string {
	if l == nil {
		l = logger.NewLogger("[CosArgs]")
	}
	value := *ArgCosVersion
	if value == "" {
		l.Fatal("No Tencent cos sdk version provided, please add it to your args use the name -bVersion")
	}
	if value != "v5" && value != "v4" {
		l.Fatal("Tencent cos sdk version provided as v5 or v4")
	}
	return *ArgCosVersion
}
