# Rikka - 极简图床系统

## 简介

Rikka（因为是罗马音，读音类似`莉卡`而不是`瑞卡`）是一套完整的个人图床系统，她包括：

- 一个 Web 界面（详见 [Demo](#Demo)）
- 一个 RESTFul API 后端（详见 [API][api-doc]）
- 基于 API 的命令行工具 rikkac（详见[CLI][cli-doc]）
- 各种实际图片的储存插件（详见[插件](#插件)）

计划实现的其他系统组件：

- Android 客户端
- iOS 客户端

Rikka 的镜像已经发布到了 [DockerHub](https://hub.docker.com/r/7sdream/rikka/), 直接开始用吧。

目前 Docker Image latest tag 版本：0.1.2

## 特点

1. 极简，不保存上传历史
2. 支持将图片链接复制成多种格式（已完成）
3. 文件储存部分插件化。（见[插件](#插件)一节）
4. 提供 API（已完成）
4. Web 服务和 RESTful API 服务模块化 （已完成）
5. CLI 工具（已完成）
6. **只对最新版 Chrome/Firefox/Safari 保持兼容**
7. 首页标志很可爱
8. 维护者很活跃（貌似……

\*：没错这是优点。如果你遇到无法在预览页面复制地址，或者一直停留在 please wait 界面，那就基本上肯定是你的浏览器不支持 es6 的锅咯。因为我前端不擅长啊，刚看了几天 es6 就被逼上阵写了点 js，实在是心有余而力不足，如果有谁愿意帮忙改善兼容性的话，辣就太蟹蟹里辣！

## Demo

这里有一个使用 Rikka 建立的[网站 Demo][demo]，密码是 `rikka`。

主页大概长这样:

![homepage][home]

点击 `Choose` 按钮选一张图片。

输入密码 `rikka`。

点击上传按钮。

上传完成后你将转到查看页面:

![viewpage][view]

如果文件过大，还没有保存完毕的话会看到等待提示，等一下就好。

等地址出现后，点击 `Src`, `Markdown`, `HTML`, `RST` 按钮可以复制对应格式的文本，然后你可以把它粘贴到其他地方。

但是注意：如果你关闭了这个页面，除了浏览器的历史记录（或者你保存了这个网址），网站并没有提供其他让你找到以前上传的图片的方法。

这是有意为之的，因为 Rikka 的主要设计的理念就是简单， `上传-复制-粘贴-关闭`，之后就再也不用管了。

PS：你看到的这些预览图也是由 Rikka 储存的哟。

## 插件

Rikka 的真实储存后端使用插件形式编写。可通过 `-plugin` 参数设置。

请看 [Rikka 插件文档][plugins-doc] 查看目前可用的插件。

## API

请看 [Rikka API 文档][api-doc]。

## CLI

Rikkac 是基于 Rikka 的 RESTful API 写的命令行工具。

编译、配置和使用方法请看 [Rikka CLI 文档][cli-doc]。

## 部署

想部署自己的 Rikka 系统？请看 [Rikka 部署文档](deploy-doc)

## License

Rikka 系统的所有代码均基于 MIT 协议开源。

详见 LICENSE 文件。

[demo]: http://7sdream-rikka-demo.daoapp.io/
[home]: http://7sdream-rikka-demo.daoapp.io/files/2016-09-04-097924191
[view]: http://7sdream-rikka-demo.daoapp.io/files/2016-09-04-017113138

[api-doc]: https://github.com/7sDream/rikka/tree/master/api
[cli-doc]: https://github.com/7sDream/rikka/tree/master/cli
[plugins-doc]: https://github.com/7sDream/rikka/tree/master/plugins
[deploy-doc]: https://github.com/7sDream/rikka/blob/master/deploy.md
