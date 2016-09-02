# Rikka - A simple photo share website.

[中文版](https://github.com/7sDream/rikka/blob/master/README.zh.md)

## Demo

I build a [Demo website](http://7sdream-rikka-demo.daoapp.io/) use rikka, it's password is `rikka`.

You can see the homepage:

![](http://7sdream-rikka-demo.daoapp.io/files/2016-09-02-544100677)

Click `Choose` button to select a file to upload.

Input rikka password.

Click `Upload` button and wait.

Then you get:

![](http://7sdream-rikka-demo.daoapp.io/files/2016-09-02-734641087)

Click `Src`, `Markdown`, `HTML`, `RST` button to copy the corresponding text to the clipboard, and paste to anywhere you want.

But, if you close this page, you have no way to find it back except from browser history(Or you save this url to other place). 

This is intentional, Because main design concept of Rikka is Simple, `Upload-Copy-Paste-Close`，then you can forget Rikka.

## Deploy

### Way 1: Build Rikka on your VPS

1. `go get github.com/7sDream/rikka`
2. `cd $GOPATH/src/github.com/7sDream/rikka`
3. `go build github.com/7sDream/rikka`
4. `./rikka --port 80 --pwd yourpassword`

Last step may require `sudo`, because Rikka use `80` port as default.

Then you can view your website and use the password you set to upload and share photo.

### Way 2: Use Docker

Docker image published to [DockerHub](https://hub.docker.com/r/7sdream/rikka/), just use it.

1. `docker pull 7sdream/rikka`
2. `docker run -d -P 7sdream/rikka:latest -pwd yourpassword`

Visit your domain or ip address with your browser and test it.

PS: If your stop/rm this container, your photo file will be deleted too. If you don't want this, use docker volume described bellow.

#### Add volume when run rikka

1. Create a vloume: `docker volume create --name rikkafiles`
2. add option `-v rikkafiles:/go/src/github.com/7sDream/rikka/files` when run rikka image

### Way 3: Use docker cloud services provider

For example, we can use DaoCloud(free qutoa) to deploy a Rikka server,

See [daocloud depoly tutorial](https://github.com/7sDream/rikka/wiki/%E5%9C%A8-DaoCloud-%E4%B8%8A%E5%85%8D%E8%B4%B9%E9%83%A8%E7%BD%B2-Rikka) for detail steps.
