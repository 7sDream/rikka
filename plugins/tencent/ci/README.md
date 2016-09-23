# Tencent CI —— Cloud Image Plugin

[中文版][version-zh]

Added in version 0.4.0, inner name `tcci`.

## Description

This plugin use Cloud Image (CI) of Tencent to store image files.

## Options

You should provide 4 options: APPID, Secret ID, Secret Key and Bucket Name.

First three options should be provided in env var, use key `RIKKA_TENCENT_APPID`, `RIKKA_TENCENT_SECRETID` and `RIKKA_TENCENT_SECRETKEY`.

And the Bucket Name should be specified by the command line option `-bname`.

If you want, you can use option `-bpath` to set the path image will be store to.

For example, `-bpath rikka`，will save image in `rikka` folder.

## Guide

See [Rikka Deploy Guide with TC-CI Plugin on DaoCloud][tcci-plugin-guide].

[version-zh]: https://github.com/7sDream/rikka/blob/master/plugins/tencent/ci/README.zh.md
[tcci-plugin-guide]: https://github.com/7sDream/rikka/wiki/%E4%BD%BF%E7%94%A8%E8%85%BE%E8%AE%AF-CI-%E6%8F%92%E4%BB%B6
