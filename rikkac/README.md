# Rikkac - CLI tool of Rikka

[中文版][version-zh]

Rikkac need to be used with a [Rikka][rikka] server.

## Usage

`rikkac <format> filename`

`<format>` can be:

- `-s`: Src, image source url
- `-m`: Markdown
- `-h`: HTML
- `-b`: BBCode
- `-r` reStructuredText

Src is default format. Format priority as same as the list ablove, lowest to highest. This is, `-m -b` considered as `-b`, `-m` is ignored. Not so complicated, you shouldn't remember priority if you never provide two format in one command.

## Build and Install

`go get github.com/7sDream/rikka/rikkac`

Add `$GOPATH/bin` into your `PATH`, if you havn't do this when you install Golang.

Then run `rikkac --version`, a version number means install successfully.

You need some configure before use Rikkac.

## Configure and Usage

Rikkac need to env variable： `RIKKA_HOST` and `RIKKA_PWD`. for  Rikka server address and password.

![](http://7sdream-rikka-demo.daoapp.io/files/2016-09-05-066558195)

Than you can enjoy Rikkac.

Just run `rikkac -m filepath` for upload, `rikkac filepath -m` is also runnable.

You can get detail log when you meet some error by add  `-v` or `-vv` option.

## Tips: Copy Result to Clipboard in Quick

![](http://7sdream-rikka-demo.daoapp.io/files/2016-09-05-781037494)

need xclip installed：`apt-get install xclip`.

[version-zh]: https://github.com/7sDream/rikka/blob/master/rikkac/README.zh.md

[rikka]: https://github.com/7sDream/rikka
