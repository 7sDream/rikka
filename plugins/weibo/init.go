package weibo

import "github.com/7sDream/rikka/common/util"

func (wbp weiboPlugin) Init() {
	l.Info("Start plugin weibo")

	cookiesStr := util.GetEnvWithCheck("Cookies", cookiesEnvKey, l)

	client = newWeiboClient()

	if err := updateCookies(cookiesStr); err != nil {
		l.Fatal("Error happened when create cookies:", err)
	}

	l.Info("Arg update cookies password =", *argUpdateCookiesPassword)

	l.Info("Weibo plugin start successfully")
}
