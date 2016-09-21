# Weibo 新浪微博插件

[English version][version-en]

0.3.1 版本添加。内部名 `weibo`.

## 简介

这个插件使用新浪微博发送微博时的上传图片接口作为图片的最终储存方式。

## 参数

你需要提供一个已登录的微博用户的 Cookies 字符串，储存在环境变量 `RIKKA_WEIBO_COOKIES` 里。

你还需要通过 `-ucpwd` 参数提供一个用于在 Web 界面（`/cookies`）上更新 Cookies 时需要输入的密码。如果你不提供的话，默认密码是 `weibo`。（`ucpwd` 的意思是 Update Cookies PassWorD）

Cookie 字符串的格式大概是：

```text
FOO=foofoofoof; BAR=barbarbarb; ZOO=zoozozozozozo
```

注意：你需要提供 weibo.com 下的**所有** Cookies，包括含有 `HTTPOnly` 属性的。

## 获取完整 Cookies 字符串

1. 启动 **Chrome** 浏览器
2. 访问 http://weibo.com
3. 登录微博（如果现在没登录的话）
4. 按 `F12` 打开开发人员工具， 转到 `Network`
5. 刷新页面
6. 找到请求列表里的第一个请求（以 `home` 开头），点击它
7. 在右边的请求内容里找到 `Request Header` 里的 `Cookies` 字段，复制字段值。（不包括前面的 `Cookies: `）

图文教程请看[部署教程](#部署教程)一节。

## 启动后更新 Cookies

在你部署并启动 Rikka 后，你可以在 Cookies 过期后通过一个 Web 页面更新它。

访问 `/cookies` 页面，把新的 Cookies 字符串复制进第一个文本框里，第二个框里填 Cookies 更新密码（就是在启动 Rikka 时提供的 `-ucpwd` 参数），点击提交。

如果更新成功，你会被转到 Rikka 的首页。如果失败了则会显示错误信息。

## 部署教程

请看部署教程：[在 DaoCloud 上部署使用新浪微博插件的 Rikka][weibo-plugin-guide]。

[version-en]: https://github.com/7sDream/rikka/blob/master/plugins/weibo/README.md
[weibo-plugin-guide]: https://github.com/7sDream/rikka/wiki/%E4%BD%BF%E7%94%A8%E6%96%B0%E6%B5%AA%E5%BE%AE%E5%8D%9A%E6%8F%92%E4%BB%B6
