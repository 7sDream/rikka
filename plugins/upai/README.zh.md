# UPai 插件

[English version][version-en]

0.3.0 版本添加。内部名 `upai`。

## 介绍

这个插件使用又拍云 CND 来储存你上传的图片。

## 参数

你需要提供四个参数：又拍云的操作员名和密码，以及图片要保存到的空间名和空间域名。

操作员名和密码使用环境变量的形式提供，变量名 `RIKKA_UPAI_OPERATOR` 和 `RIKKA_UPAI_PASSWORD`。

空间名和空间域名则通过命令行参数提供：

`-bname` 空间名

`-bhost` 空间域名

另外，你还可以通过提供 `bpath` 参数的形式设置图片需要保存到的文件夹。

比如，使用 `-bpath rikka`，上传到的文件会传到空间的 `rikka` 文件夹下。当然 `-bpath rikka/images` 之类的多级文件夹也是可以的。

## 部署教程

正在编写中。

[version-en]: https://github.com/7sDream/rikka/blob/master/plugins/upai/README.md
