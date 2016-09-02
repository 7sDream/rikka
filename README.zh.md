# Rikka - 极简图床

[English version](https://github.com/7sDream/rikka)

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

最后一步因为要使用 80 端口，所以可能需要 `sudo`。

之后你就可以用浏览器打开看看效果了。

### 方式 2: 使用 Docker

Rikka 的镜像已经发布到了 [DockerHub](https://hub.docker.com/r/7sdream/rikka/), 直接开始用吧。

1. `docker pull 7sdream/rikka`
2. `docker run -d -P 7sdream/rikka:latest -pwd yourpassword`

打开浏览器访问你的 IP 或域名试用看看吧。

PS: 如果你停止/删除了 Rikka 容器，你上传的照片也会一起被删除。如果你不想这样，请参考下一节：使用数据卷。

#### 使用数据卷

Docker 提供了数据卷的功能，这样就不用爬和 Rikka 无关我们上传的图片在应用关闭之后丢失了。

使用方法：

1. 创建数据卷：`docker volume create --name rikkafiles`
2. 在启动 Rikka 容器时加上如下参数：`-v rikkafiles:/go/src/github.com/7sDream/rikka/files`

### 方式 3: 使用 Docker 云服务提供商

比如，我们可以用 DaoCloud 的免费配额来部署一个 Rikka 服务。

详细步骤请看 [DaoCloud 部署教程](https://github.com/7sDream/rikka/wiki/%E5%9C%A8-DaoCloud-%E4%B8%8A%E5%85%8D%E8%B4%B9%E9%83%A8%E7%BD%B2-Rikka)。
