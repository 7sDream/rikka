# Rikka Plugin System

[中文版][version-zh]

Rikka back-end image save is powered by plugins.

## fs Plugin

This is default plugin of Rikka，it save image file to server where Rikka live directly, and run a static file server for those files.

Refer to [fs Plugin Doc][fs-doc] for plugin options.

## Qiniu Cloud Plugin

This plugin use Qiniu Cloud CDN to store your image.

Refer to [Qiniu Plugin Doc][qiniu-doc] for plugin options.

## UPai Cloud Plugin

This plugin use UPai Cloud CDN to store your image.

Refer to [UPai Plugin Doc][upai-doc].

## Sina Weibo Plugin

This plugin use Sina weibo to store your image.

Refer to [Weibo plugin Doc][weibo-doc].

## Tencent COS Plugin

This plugin use Cloud Object Service (COS) of Tencent to store image files.

Refer to [Tccos Plugin Doc][tccos-doc].

[version-zh]: https://github.com/7sDream/rikka/blob/master/plugins/README.zh.md

[fs-doc]: https://github.com/7sDream/rikka/tree/master/plugins/fs
[qiniu-doc]: https://github.com/7sDream/rikka/tree/master/plugins/qiniu
[upai-doc]: https://github.com/7sDream/rikka/tree/master/plugins/upai
[weibo-doc]: https://github.com/7sDream/rikka/tree/master/plugins/weibo
[tccos-doc]: https://github.com/7sDream/rikka/tree/master/plugins/tencent/cos
