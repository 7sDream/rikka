# Rikka - 极简图床系统

![][badge-version-img] ![][badge-info-img] ![][badge-license-img]

[English version][readme-en]

Rikka 主要使用 Go 语言编写，并提供 Docker 镜像。

Rikka 的镜像已经发布到了 [DockerHub][image-in-docker-hub], 直接开始用吧。

最新版本号和镜像大小见上面的徽章。

## 简介

Rikka（因为是日文罗马音，读音类似`莉卡`而不是`瑞卡`）是一套完整的个人图床系统，她包括：

- 一个 Web 应用（详见 [Demo](#demo) 一节）
- 一个 REST API 后端（详见 [API 文档][api-doc]）
- 基于 API 的命令行工具 Rikkac（详见 [Rikkac 文档][rikkac-doc]）
- 图片的实际储存插件（查看[插件文档][plugins-doc] 来获取所有可用插件的列表）

计划实现的其他非 Go 语言的系统组件：

- Android 客户端
- iOS 客户端

## 特点

1. 极简，不保存上传历史
2. 支持将图片链接复制成多种格式
3. 文件储存部分插件化，有很多可用的插件，比如：新浪微博，七牛云，又拍云，腾讯云等
4. 提供 API
4. Web 服务和 REST API 服务模块化
5. CLI 工具
6. **只保证支持较新版本的 Chrome/Firefox/Safari**
7. 首页标志很可爱
8. 维护者貌似很活跃 ：）

## Demo

这里有一个使用 Rikka 建立的[网站 Demo][demo]，~~密码是 `rikka`~~，由于 DaoCloud 现在不能免费用了，所以现在这里的 demo 其实是我自己用的，所以大家只能看看主页了。

主页大概长这样:

![homepage][home]

点击 `Choose` 按钮选一张图片。

输入密码 `rikka`。

点击上传按钮。

上传完成后你将转到查看页面:

![view_page][view]

如果文件过大，还没有保存完毕的话会看到等待提示，等一下就好。

等地址出现后，点击 `Src`, `Markdown`, `HTML`, `RST`，`BBCode` 按钮可以复制对应格式的文本，然后你可以把它粘贴到其他地方。

但是注意：如果你关闭了这个页面，除了浏览器的历史记录（或者你保存了这个网址），网站并没有提供其他让你找到以前上传的图片的方法。

这是有意为之的，因为 Rikka 的主要设计的理念就是简单， `上传-复制-关闭-粘贴`，之后就再也不用管了。

PS：你看到的这些预览图也是由 Rikka 储存的哟。（不过放到 Github 之后会被 Github 弄到 CDN 上去）

## 插件

Rikka 的真实储存后端使用插件形式编写。可通过 `-plugin` 参数设置。

请看 [Rikka 插件文档][plugins-doc] 查看目前可用的插件。

## API

请看 [Rikka API 文档][api-doc]。

## CLI - Rikkac

Rikkac 是基于 Rikka 的 REST API 写的 Rikka CLI 工具。

编译、安装、配置和使用方法请看 [Rikkac 文档][rikkac-doc]。

## 部署

想部署自己的 Rikka 系统？请看 [Rikka 部署文档][deploy-doc]。

## 协助开发

- Fork
- 从 dev 分支新建一个分支
- 写代码，注释和文档，并使用有意义的 commit message 提交
- 将自己加入 CONTRIBUTIONS.md，并且描述你做了什么
- PR 到 dev 分支

感谢所有协助开发的朋友！

在 [CONTRIBUTIONS.md][contributors] 里可以看到贡献者名单。

## 致谢

- 感谢 Go 编程语言以及她的开发团队
- 感谢 Visual Studio Code 编辑器和她的开发团队
- 感谢开源精神

## License

Rikka 系统的所有代码均基于 MIT 协议开源。

详见 [LICENSE][license] 文件。

[readme-en]: https://github.com/7sDream/rikka/blob/master/README.md

[badge-info-img]: https://images.microbadger.com/badges/image/7sdream/rikka.svg
[badge-version-img]: https://images.microbadger.com/badges/version/7sdream/rikka.svg
[badge-license-img]: https://images.microbadger.com/badges/license/7sdream/rikka.svg

[image-in-docker-hub]: https://hub.docker.com/r/7sdream/rikka/

[demo]: https://rikka.7sdre.am/
[home]: https://rikka.7sdre.am/files/56c3ae9d-4d96-49c8-bc03-5104214a1ac8.png
[view]: https://rikka.7sdre.am/files/97bebf3b-9fb8-4b0c-a156-4b92b1951ae4.png

[api-doc]: https://github.com/7sDream/rikka/blob/master/api/README.zh.md
[rikkac-doc]: https://github.com/7sDream/rikka/blob/master/rikkac/README.zh.md
[plugins-doc]: https://github.com/7sDream/rikka/blob/master/plugins/README.zh.md
[deploy-doc]: https://github.com/7sDream/rikka/blob/master/deploy.zh.md

[contributors]: https://github.com/7sDream/rikka/blob/master/CONTRIBUTORS.md
[license]: https://github.com/7sDream/rikka/blob/master/LICENSE
