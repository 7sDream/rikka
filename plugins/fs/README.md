# fs Plugin

Inner name `fs`， default plugin of Rikka.

## Description

it save image file to server where Rikka live directly, and run a static file server for those files.

## Options

`-dir` set file dir where image saved。Default is `files` folder under workdir. If you are using Docker or deploying Rikka at Docker Cloud Server Provider, you can set it to a position easy to volume mount, like `/data`.

`-fsDebugSleep` Not for common use, it make a sleep before copy file to dir, simulate a long time operation，for javascript AJAX tests. In microsecond.
