# Qiniu Plugin

[中文版][version-zh]

Added in version 0.2.0. Inner name `qiniu`.

## Description

This plugin use Qiniu Cloud CND to store your image.

## Options

You should provide Qiniu ACCESS KEY, SECRET KEY, bucket name and bucket host.

ACCESS KEY and SECRET KEY should be add into your env variable, use key `RIKKA_QINIU_ACCESS` and `RIKKA_QINIU_SECRET`.

Bucket name and bucket host should be provide use command line option:

`-bname` for the bucket name.

`-bhost` for bucket host.

BTW： you can set upload dir by provide `-bpath` option.

For example，ues `-bpath rikka`, then images will be under `rikka` folder。

Multi-level dir like `-bpath rikka/images` are also supported.

## Guide

See [Rikka Deploy Guide with Qiniu Plugin on DaoCloud][qiniu-plugin-guide].

[version-zh]: https://github.com/7sDream/rikka/blob/master/plugins/qiniu/README.zh.md
[qiniu-plugin-guide]: https://github.com/7sDream/rikka/wiki/%E4%BD%BF%E7%94%A8%E4%B8%83%E7%89%9B%E4%BA%91%E6%8F%92%E4%BB%B6
