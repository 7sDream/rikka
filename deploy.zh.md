# 部署

[English version][version-en]

以下部署方法均以默认 `fs` 插件为例。

## 方式 1: 在你的 VPS 上编译

1. `go get -u -d github.com/7sDream/rikka`
2. `cd $GOPATH/src/github.com/7sDream/rikka`
3. `go build .`
4. `./rikka -port 80 -pwd yourPassword`

最后一步具体的命令可查看 `./rikka -h` 之后根据自己需要设置。

因为要使用 80 端口，所以可能需要在启动命令前加上 `sudo`。

之后你就可以用浏览器打开看看效果了。

## 方式 2: 使用 Docker

1. `docker pull 7sdream/rikka`
2. `docker run -d -p 80:80 7sdream/rikka -pwd yourPassword`

同样可以根据需要设定参数。至于 image expose 的是 80 端口，请根据需要进行映射。 

打开浏览器访问你的 IP 或域名试用看看吧。

PS: 如果你停止/删除了 Rikka 容器，你上传的照片也会一起被删除。如果你不想这样，请参考下一节：使用数据卷。

### 使用数据卷

Docker 提供了数据卷的功能，这样就不用怕我们上传的图片会应用关闭之后丢失了。

使用方法：

1. 创建数据卷：`docker volume create --name rikka_files`
2. 在启动 Rikka 容器时加上如下参数：`-v rikka_files:/go/src/github.com/7sDream/rikka/files`

PS：你可以使用 Rikka `fs` 插件的 `-dir` 参数指定文件储存位置，比如这样：

`docker run -d -P -v rikka_files:/data --name rikka 7sdream/rikka -pwd 12345 -dir /data`

这样就不用把挂载路径设的太长了。

## 方式 3: 使用 Docker 云服务提供商

比如，我们可以用 DaoCloud 的免费配额来部署一个 Rikka 服务。

详细步骤请看 [DaoCloud 部署教程](https://github.com/7sDream/rikka/wiki/%E5%9C%A8-DaoCloud-%E4%B8%8A%E5%85%8D%E8%B4%B9%E9%83%A8%E7%BD%B2-Rikka)。

## 使用其他插件

主要步骤和上述相同。

不同插件的不同启动参数请参考[插件文档][plugins-doc]。

[version-en]: https://github.com/7sDream/rikka/blob/master/deploy.md

[daocloud-guide]: https://github.com/7sDream/rikka/wiki/%E5%9C%A8-DaoCloud-%E4%B8%8A%E5%85%8D%E8%B4%B9%E9%83%A8%E7%BD%B2-Rikka
[plugins-doc]: https://github.com/7sDream/rikka/blob/master/plugins/README.zh.md
