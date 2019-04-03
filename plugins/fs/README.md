# fs Plugin

[中文版][version-zh]

Inner name `fs`， default plugin of Rikka.

## Description

it save image file to server where Rikka live directly, and run a static file server for those files.

## Options

`-dir` set file dir where image saved. Default is `files` folder under work dir. If you are using Docker or deploying Rikka at Docker Cloud Server Provider, you can set it to a position easy to volume mount, like `/data`.

`-fsDebugSleep` Not for common use, it make a sleep before copy file to dir, simulate a long time operation，for javascript AJAX tests. In microsecond.

If your website support https，you can add `-https` argument to make fs plugin return image url with https protocol.

[version-zh]: https://github.com/7sDream/rikka/blob/master/plugins/fs/README.zh.md
