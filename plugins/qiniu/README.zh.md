# Qiniu 插件

[English version][version-en]

内部名 `qiniu`。

## 介绍

这个插件使用七牛云 CND 来储存你上传的图片。

## 参数

你需要提供四个参数：七牛的 `ACCESSKEY`, `SECRETKEY`， 以及图片要保存到的空间名和空间域名。

`ACCESSKEY` 和 `SECRETKEY` 使用环境变量的形式提供，变量名 `RIKKA_QINIU_ACCESS` 和 `RIKKA_QINIU_SECRET`。

空间名和空间域名则通过命令行参数提供：

`-bname` 空间名

`-bhost` 空间域名

## 部署教程

请看部署教程：[在 DaoCloud 上部署使用七牛云插件的 Rikka][qiniu-plugin-guide]。

[version-en]: https://github.com/7sDream/rikka/blob/master/plugins/qiniu/README.md
[qiniu-plugin-guide]: https://github.com/7sDream/rikka/wiki/%E4%BD%BF%E7%94%A8%E4%B8%83%E7%89%9B%E4%BA%91%E6%8F%92%E4%BB%B6
