# Rikka-CLI - Rikka 的命令行工具

需要和 [Rikka](https://github.com/7sDream/rikka) 配合使用。

## 使用方式

`rikkac -smhbsr filename`

- `-s`: SRC 图片原始地址
- `-m`: Markdown 格式
- `h`: HTML 格式
- `b`: BBCode 格式
- `r` reStructuredText 格式

默认是源地址格式，优先级如上表，从低到高。（也就是说下面的会覆盖上面的，`-m -b` 等同于 `-b`）

## 截图


