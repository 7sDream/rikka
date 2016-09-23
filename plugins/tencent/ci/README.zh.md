# TC-CI 腾讯万象优图 Cloud Image 插件

[English version][version-en]

0.4.0 版本添加，内部名 `tcci`。

## 简介

这个插件使用腾讯云的万象优图（Colud Image, CI）来储存图片。

## 参数

你需要提供四个参数：腾讯云的项目编号（APPID），密钥ID（SecretID），密钥Key（SecretKey）以及储存空间名（bucket name）。

前三个通过环境变量提供，分别为 `RIKKA_TENCENT_APPID`, `RIKKA_TENCENT_SECRETID`, `RIKKA_TENCENT_SECRETKEY`。

储存空间名通过命令行参数 `-bname` 提供。

另外，你还可以通过提供 `bpath` 参数的形式设置图片需要保存到的文件夹。

比如，使用 `-bpath rikka`，上传到的文件会传到空间的 `rikka` 文件夹下。

## 部署教程

请看部署教程：[在 DaoCloud 上部署使用 TC-CI 插件的 Rikka][tcci-plugin-guide]

[version-en]: https://github.com/7sDream/rikka/blob/master/plugins/tencent/ci/README.md
[tcci-plugin-guide]: https://github.com/7sDream/rikka/wiki/%E4%BD%BF%E7%94%A8%E8%85%BE%E8%AE%AF-CI-%E6%8F%92%E4%BB%B6
