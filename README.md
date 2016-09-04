# Rikka - 极简图床

Rikka 的镜像已经发布到了 [DockerHub](https://hub.docker.com/r/7sdream/rikka/), 直接开始用吧。

目前 Docker Image latest tag 版本：0.1.0

## 特点

1. 极简，不保存上传历史
2. 支持将图片链接复制成多种格式
3. 文件储存部分插件化。虽然目前只有一个插件是将文件储存到本机，但是慢慢我会加一些比如七牛云 CDN 的插件（计划中
4. Web 服务和 RESTful API 服务模块化 
5. CLI 工具（计划中
6. **只对最新版 Chrome 保持兼容**（没错这是优点）
7. 首页标志很可爱
8. 维护者很活跃（貌似……

## 启动参数

`-bind` 指定监听的 IP 地址，默认不填的话是监听所有 IP。

`-port` 是端口，默认 80，不用多说了。

`-plugin` 设置文件储存的后端插件，现在只有一个插件 `fs`，表示直接储存在机器中。

`-dir` 参数是 `fs` 插件的参数，指定文件存放位置。如果你在 Docler 云服务上部署的话，可以设置成 `/data` 之类便于挂载的位置。

`-pwd` 参数指定上传文件时的密码。

`-size` 指定允许上传的最大文件大小，以 MB 为单位，可以有小数。

`-level` 设置日志级别

## Demo

这里有一个使用 Rikka 建立的[网站 Demo](http://7sdream-rikka-demo.daoapp.io/)，密码是 `rikka`。

主页大概长这样:

![](http://7sdream-rikka-demo.daoapp.io/files/2016-09-02-544100677)

点击 `Choose` 按钮选一张图片。

输入密码 `rikka`。

点击上传按钮。

上传完成后你将转到查看页面:

![](http://7sdream-rikka-demo.daoapp.io/files/2016-09-02-734641087)

点击 `Src`, `Markdown`, `HTML`, `RST` 按钮可以复制对应格式的文本，然后你可以把它粘贴到其他地方。

但是注意：如果你关闭了这个页面，除了浏览器的历史记录（或者你保存了这个网址），网站并没有提供其他让你找到以前上传的图片的方法。

这是有意为之的，因为 Rikka 的主要设计的理念就是简单， `上传-复制-粘贴-关闭`，之后就再也不用管了。

## 部署

### 方式 1: 在你的 VPS 上编译

1. `go get github.com/7sDream/rikka`
2. `cd $GOPATH/src/github.com/7sDream/rikka`
3. `go build github.com/7sDream/rikka`
4. `./rikka --port 80 --pwd yourpassword`

最后一步具体的命令可查看 `./rikka -h` 之后根据自己需要设置。因为要使用 80 端口，所以可能需要 `sudo`。

之后你就可以用浏览器打开看看效果了。

### 方式 2: 使用 Docker

1. `docker pull 7sdream/rikka`
2. `docker run -d -P 7sdream/rikka:latest -pwd yourpassword`

同样可以根据需要设定参数。

打开浏览器访问你的 IP 或域名试用看看吧。

PS: 如果你停止/删除了 Rikka 容器，你上传的照片也会一起被删除。如果你不想这样，请参考下一节：使用数据卷。

#### 使用数据卷

Docker 提供了数据卷的功能，这样就不用爬和 Rikka 无关我们上传的图片在应用关闭之后丢失了。

使用方法：

1. 创建数据卷：`docker volume create --name rikkafiles`
2. 在启动 Rikka 容器时加上如下参数：`-v rikkafiles:/go/src/github.com/7sDream/rikka/files`

PS：你可以使用 Rikka 的 `-dir` 参数指定文件储存位置，比如这样：

`docker run -d -P -v rikkafiles:/data --name rikka 7sdream/rikka:latest -pwd 12345 -dir /data`

这样就不用把挂载路径设的太长了。

### 方式 3: 使用 Docker 云服务提供商

比如，我们可以用 DaoCloud 的免费配额来部署一个 Rikka 服务。

详细步骤请看 [DaoCloud 部署教程](https://github.com/7sDream/rikka/wiki/%E5%9C%A8-DaoCloud-%E4%B8%8A%E5%85%8D%E8%B4%B9%E9%83%A8%E7%BD%B2-Rikka)。

## API

请看 [Rikka API 文档](https://github.com/7sDream/rikka/blob/master/API.md)
