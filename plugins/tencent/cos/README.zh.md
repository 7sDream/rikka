# TC-COS 腾讯对象储存服务 Cloud Object Service 插件

[English version][version-en]

0.4.0 版本添加，内部名 `tccos`。

## 简介

这个插件使用腾讯云的对象储存服务（Cloud Object Service, COS）来储存图片。

## 参数

你需要提供四个参数：腾讯云的项目编号（APPID），密钥ID（SecretID），密钥Key（SecretKey）以及储存空间名（bucket name）。

前三个通过环境变量提供，分别为 `RIKKA_TENCENT_APPID`, `RIKKA_TENCENT_SECRETID`, `RIKKA_TENCENT_SECRETKEY`。

储存空间名通过命令行参数 `-bname` 提供。

另外，你还可以通过提供 `bpath` 参数的形式设置图片需要保存到的文件夹。（注意，文件夹必须在 COS 里已经存在）

比如，使用 `-bpath rikka`，上传到的文件会传到空间的 `rikka` 文件夹下。

对象存储版本通过命令行参数 `-tccosVer` 提供，默认版本为 v4, v5 版本需要设置域名中的所属地域(region),添加环境变量 `RIKKA_TENCENT_REGION`

## 注意事项

腾讯 COS 作为「文件储存服务」，其默认对于存放的文件提供的是下载服务而非预览服务。

也就是说如果你把一个储存在 COS 上的图片链接直接用浏览器打开，浏览器会执行下载动作而不是在标签页里预览图片。

但是放在 HTML 的 `img` 元素的 `src` 属性里，或者其他元素的 `background` 属性里作为图片显示是没有问题的。

如果想打开链接时触发图片预览，请在 COS 设置页面绑定自定义域名并打开静态网站选项。

参考：[腾讯云 COS 静态网站选项][tencent-cos-static-website-doc]。

## 部署教程

请看部署教程：[在 DaoCloud 上部署使用 TC-COS 插件的 Rikka][tccos-plugin-guide]

[version-en]: https://github.com/7sDream/rikka/blob/master/plugins/tencent/cos/README.md
[tencent-cos-static-website-doc]: https://www.qcloud.com/doc/product/227/%E9%85%8D%E7%BD%AE%E8%AF%A6%E6%83%85#5-静态网站
[tccos-plugin-guide]: https://github.com/7sDream/rikka/wiki/%E4%BD%BF%E7%94%A8%E8%85%BE%E8%AE%AF-COS-%E6%8F%92%E4%BB%B6
