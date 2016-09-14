# Rikkac - Rikka 的命令行工具

[English version][version-en]

需要和 [Rikka][rikka] 配合使用。

## 使用方式

`rikkac <format> filename`

`<format>` 可选的参数如下:

- `-s`: SRC 图片原始地址
- `-m`: Markdown 格式
- `-h`: HTML 格式
- `-b`: BBCode 格式
- `-r` reStructuredText 格式

默认是源地址格式，优先级如上表，从低到高。也就是说下面的会覆盖上面的，`-m -b` 等同于 `-b`。其实也没那么复杂，你只要不同时提供两个就不用记优先级。

## 编译安装

### 下载二进制文件

目前编译好的 Rikkac 工具只提供了 [Linux 版下载][download]，因为我这里只有 Linux 系统 QwQ

下载了之后重命名为 `rikkac`，放到某个 `PATH` 目录下即可。

使用其他操作系统的用户请使用下一节所说的从源代码安装。

### 从源代码安装

首先你需要安装 Go，然后：

`go get github.com/7sDream/rikka/rikkac`

把 `$GOPATH/bin` 加入 `PATH` 如果你在安装 Go 的时候没做这步的话。

然后输入 `rikkac --version` 如果输出了一个版本号则说明安装成功了。

编译和安装成功后并不能立即使用，需要进行一些配置。

## 配置和使用

Rikkac 需要配置两个环境变量： `RIKKA_HOST` 和 `RIKKA_PWD`。它们分别代表 Rikka 服务器地址和密码。

![](http://7sdream-rikka-demo.daoapp.io/files/2016-09-05-066558195)

配置完就可以使用啦。

基本上就是 `rikkac -m filepath` 就好。

如果出错了可以用 `-v` 或者 `-vv` 参数输出详细日志用于排错。

## 批量上传

`rikkac -m file1 file2 file3 ...` 这样就行了。

如果你用的 shell 带有通配符自动展开的话，那这样也行：`rikkac -m *.png`。

![](http://odbw8jckg.bkt.clouddn.com/ba2d2dca-2ae2-4436-ade2-7905183ce23d.png)

## 小 tipc 快速复制到剪贴板

![](http://7sdream-rikka-demo.daoapp.io/files/2016-09-05-781037494)

此方法需要安装 xclip：`apt-get install xclip`。

[version-en]: https://github.com/7sDream/rikka/blob/master/rikkac/README.md

[rikka]: https://github.com/7sDream/rikka/blob/master/README.zh.md
[download]: https://github.com/7sDream/rikka/releases/tag/Rikkac
