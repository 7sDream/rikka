# Qiniu 插件

内部名 `qiniu`。

## 介绍

这个插件使用七牛云 CND 来储存你上传的图片。

## 参数

你需要提供四个参数：七牛的 `ACCESSKEY`, `SECRETKEY`， 以及图片要保存到的空间名和空间域名.

`ACCESSKEY` 和 `SECRETKEY` 使用环境变量的形式提供，变量名 `RIKKA_QINIU_ACCESS` 和 `RIKKA_QINIU_SECRET`。

空间名和空间域名则通过命令行参数提供：

`-bname` 空间名

`-bhost` 空间域名
