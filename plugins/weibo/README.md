# Sina Weibo Plugin

[中文版][version-zh]

Added in version 0.3.1. Inner name `weibo`.

## Description

This plugin use Sina weibo to store your image.

## Options

You should provide a cookies string which stand for a logged weibo account in env var `RIKKA_WEIBO_COOKIRS`.

And you should provide a password (with option `-ucpwd`, update cookies password) which will be checked when you update cookies from `/cookies` page. Default password is `weibo`.

Format of cookies string:

```text
FOO=foofoofoof; BAR=barbarbarb; ZOO=zoozozozozozo
```

Notice: You should provide **ALL** cookies of weibo.com, contains thoose be tag with **HTTPOnly**.

## A way of get cookies string

1. Launch **Chrome**
2. visit http://weibo.com
3. Login if you haven't
4. `F12` to open devtools, turn to `Console` tab, type `document.cookie`, and copy output string(exclude around `""`) to a temp text file
6. Turn to `Application`(Or `Resource`) tab, click `Cookies` in left sidebar, and find `weibo.com`
7. Click `HTTPOnly` field of right table
8. Add **ALL** lines whose `HTTPOnly` field is checked in format `Name=Value; ` to the end of temp text file
9. Now the content of temp text file is your cookies string
10. Use command `export RIKKA_WEIBO_COOKIRS="<temp file content>"` to set env var

Tutorial with image can be find in [Guide](#Guide) section.

## Update cookies after launch

After you deploy and launch Rikka, you can update weibo cookies when expired.

Visit `/cookies` page, put new cookies string into first textarea, your `ucpwd` into second, and submit.

If update successfully, you will be redirect to homepage of Rikka. And Error message will shown if error happened.

## Guide

WIP.

[version-zh]: https://github.com/7sDream/rikka/blob/master/plugins/weibo/README.zh.md
