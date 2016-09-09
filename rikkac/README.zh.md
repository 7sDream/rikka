# Rikkac - Rikka 的命令行工具

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

`go get github.com/7sDream/rikka/rikkac`

把 `$GOPATH/bin` 加入 `PATH` 如果你在安装 Go 的时候没做这步的话。

然后输入 `rikkac --version` 如果输出了一个版本号则说明安装成功了。

编译和安装成功后并不能立即使用，需要进行一些配置。

## 配置和使用

Rikkac 需要配置两个环境变量： `RIKKA_HOST` 和 `RIKKA_PWD`。它们分别代表 Rikka 服务器地址和密码。

![](http://7sdream-rikka-demo.daoapp.io/files/2016-09-05-066558195)

配置完就可以使用啦。

基本上就是 `rikkac -m filepath` 就好，当然 `rikkac filepath -m` 也是可以的。

如果出错了可以用 `-v` 或者 `-vv` 参数输出详细日志用于排错。

## 小 tipc 快速复制到剪贴板

![](http://7sdream-rikka-demo.daoapp.io/files/2016-09-05-781037494)

此方法需要安装 xclip：`apt-get install xclip`。

[rikka]: https://github.com/7sDream/rikka/blob/master/README.zh.md
