# fs 插件

[English version][version-en]

插件内部名 `fs`，是 Rikka 的默认插件。

## 说明

它直接将上传的图片储存在部署 Rikka 的服务器上，并且使用一个静态文件 Server 来提供这些图片。

## 参数

`-dir` 参数指定文件存放位置。默认位置是当前目录下的 `files` 文件夹。如果你使用 Docker 或在 Docker 云服务上部署的话，可以设置成 `/data` 之类便于挂载的位置。

`-fsDebugSleep` 一般用不到，是让 fs 插件在复制文件前暂停一段时间，模拟耗时操作，便于测试 javascript AJAX 的。单位是 ms。

如果你的域名支持 https，请设置参数 `-https` 来使 fs 模块返回的 https 协议的 url。

[version-en]: https://github.com/7sDream/rikka/blob/master/plugins/fs/README.md
