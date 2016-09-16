package weibo

import "os"

func (wbp weiboPlugin) Init() {
	l.Info("Start plugin weibo")

	cookiesStr := os.Getenv(cookiesEnvKey)
	if cookiesStr == "" {
		l.Fatal("No weibo cookies provided， plesae add it into your env var use the name", cookiesEnvKey)
	}

	client = newWeiboClient()

	if err := updateCookies(cookiesStr); err != nil {
		l.Fatal("Error happened when create cookies:", err)
	}

	l.Info("Weibo plugin start successfilly")
}
