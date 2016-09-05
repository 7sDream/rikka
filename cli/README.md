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

- `go build -o rikkac github.com/7sDream/rikka/cli`
- `mv rikkac /some/dir/in/your/path`

## 配置

需要两个环境变量 `RIKKA_HOST` 和 `RIKKA_PWD`，分别为 Rikka 服务器的地址和密码。

当然，如果你愿意，也可以在使用时用 `-t` 和 `-p` 参数指定，它们的优先级比环境变量高。

## 使用

![][usage-preview]

基本上就是 `rikkac -m filepath` 就好。

如果出错了可以用 `-v` 或者 `-vv` 参数输出详细日志用于排错。

[rikka]: https://github.com/7sDream/rikka
[usage-preview]: http://7sdream-rikka-demo.daoapp.io/files/2016-09-04-221897650
