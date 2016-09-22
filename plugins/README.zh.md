# Rikka 插件系统

[English version][version-en]

Rikka 后端实际图片的储存使用插件形式处理。

## fs 插件

这是 Rikka 的默认插件，它直接将上传的图片储存在部署 Rikka 的服务器上，并且使用一个静态文件 Server 来提供这些图片。

请看 [fs 插件文档][fs-doc] 查看插件的配置参数。

## Qiniu 七牛云插件

这个插件使用七牛云 CDN 来储存你上传的图片。

请看 [Qiniu 插件文档][qiniu-doc] 查看插件的配置参数。

## Upai 又拍云插件

这个插件使用又拍云 CDN 来储存你上传的图片。

请看 [UPai 插件文档][upai-doc] 查看插件配置参数。

## weibo 新浪微博插件

这个插件使用新浪微博发送微博时的上传图片接口作为图片的最终储存方式。

请看 [Weibo 插件文档][weibo-doc] 查看插件配置参数。

[version-en]: https://github.com/7sDream/rikka/blob/master/plugins/README.md

## Tccos 腾讯 COS 插件

这个插件使用腾讯云的对象储存服务（COS）来储存图片。

请看 [Tccos 插件文档][tccos-doc] 查看插件配置参数。

[fs-doc]: https://github.com/7sDream/rikka/tree/master/plugins/fs/README.zh.md
[qiniu-doc]: https://github.com/7sDream/rikka/tree/master/plugins/qiniu/README.zh.md
[upai-doc]: https://github.com/7sDream/rikka/tree/master/plugins/upai/README.zh.md
[weibo-doc]: https://github.com/7sDream/rikka/tree/master/plugins/weibo/README.zh.md
[tccos-doc]: https://github.com/7sDream/rikka/tree/master/plugins/tencent/cos/README.zh.md
