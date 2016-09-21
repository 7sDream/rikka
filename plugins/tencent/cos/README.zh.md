# 腾讯 COS —— Cloud Object Service 插件

[English version][version-en]

0.4.0 版本添加，内部名 `tccos`。

## 简介

这个插件使用腾讯云的对象储存服务（COS）来储存图片。

## 参数

你需要提供四个参数：腾讯云的项目编号（APPID），密钥ID（SecretID），密钥Key（SecretKey）以及储存空间名（bucket name）。

前三个通过环境变量提供，分别为 `RIKKA_TCCOS_APPID`, `RIKKA_TCCOS_SECRETID`, `RIKKA_TCCOS_SECRETKEY`。

储存空间名通过命令行参数 `-bname` 提供。

另外，你还可以通过提供 `bpath` 参数的形式设置图片需要保存到的文件夹。（注意，文件夹必须在 COS 里已经存在）

比如，使用 `-bpath rikka`，上传到的文件会传到空间的 `rikka` 文件夹下。

## 部署教程

WIP

[version-en]: https://github.com/7sDream/rikka/blob/master/plugins/tencent/cos/README.md
