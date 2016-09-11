# Rikka 插件系统

[English version][version-en]

Rikka 后端实际图片的储存使用插件形式处理。

## fs 插件

这是 Rikka 的默认插件，它直接将上传的图片储存在部署 Rikka 的服务器上，并且使用一个静态文件 Server 来提供这些图片。

请看 [fs 插件文档][fs-doc] 查看插件的配置参数。

## Qiniu 七牛云插件

这个插件使用七牛云 CND 来储存你上传的图片。

请看 [Qiniu 插件文档][qiniu-doc] 查看插件的配置参数。

## weibo 新狼微博插件

也还没写出来。

[version-en]: https://github.com/7sDream/rikka/blob/master/plugins/README.md

[fs-doc]: https://github.com/7sDream/rikka/tree/master/plugins/fs/README.zh.md
[qiniu-doc]: https://github.com/7sDream/rikka/tree/master/plugins/qiniu/README.zh.md
