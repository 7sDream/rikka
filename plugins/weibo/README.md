# Sina Weibo Plugin

[中文版][version-zh]

Added in version 0.3.1. Inner name `weibo`.

## Description

This plugin use Sina weibo to store your image.

## Options

You should provide a cookies string which stand for a logged weibo account in env var `RIKKA_WEIBO_COOKIES`.

And you should provide a password (with option `-ucpwd`, update cookies password) which will be checked when you update cookies from `/cookies` page. Default password is `weibo`.

Format of cookies string:

```text
FOO=foofoofoof; BAR=barbarbarb; ZOO=zoozozozozozo
```

Notice: You should provide **ALL** cookies of weibo.com, contains those be tag with **HTTPOnly**.

## A way of get cookies string

1. Launch **Chrome**
2. visit http://weibo.com
3. Login if you haven't
4. `F12` to open dev tools, turn to `Network` tab
5. Refresh page
6. Click first request(starts with `home`) in the left list
7. Find `Cookies` field of `Request Header` in the request content(right side), copy field value(without the `Cookies: ` prefix)

Tutorial with image can be find in [Guide](#Guide) section.

## Update cookies after launch

After you deploy and launch Rikka, you can update weibo cookies when expired.

Visit `/cookies` page, put new cookies string into first textarea, your `ucpwd` into second, and submit.

If update successfully, you will be redirect to homepage of Rikka. And Error message will shown if error happened.

## Guide

See [Rikka Deploy Guide with Weibo plugin on DaoCloud][weibo-plugin-guide]

[version-zh]: https://github.com/7sDream/rikka/blob/master/plugins/weibo/README.zh.md
[weibo-plugin-guide]: https://github.com/7sDream/rikka/wiki/%E4%BD%BF%E7%94%A8%E6%96%B0%E6%B5%AA%E5%BE%AE%E5%8D%9A%E6%8F%92%E4%BB%B6
