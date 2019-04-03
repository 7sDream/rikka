# Deploy

[中文版][version-zh]

The following ways use default plugin `fs` as example.

## Way 1: Build in you VPS

1. `go get -u -d github.com/7sDream/rikka`
2. `cd $GOPATH/src/github.com/7sDream/rikka`
3. `go build .`
4. `./rikka -port 80 -pwd yourPassword`

You can use `./rikka --help` to get more options and make you own launch command.

Because port 80 wll be used, may you need `sudo` prefix.

Then you can open your browser to test Rikka.

## Way 2: Use Docker

1. `docker pull 7sdream/rikka`
2. `docker run -d -p 80:80 7sdream/rikka -pwd yourPassword`

You can set option based on you requirements. 

Rikka image expose 80 port, you can map it based on needs.

Then you can open your browser to test Rikka.

Note: If you stop/remove Rikka container, the images you uploaded will be deleted too. If you want keep those files, please read next section: Use Volume.

### Use Volume

Docker provide a feature called Volume. We can use it to keep out images.

Usage：

1. Create volume：`docker volume create --name rikka_files`
2. Add this option when you start Rikka：`-v rikka_files:/go/src/github.com/7sDream/rikka/files`

BTW: You can use `-dir` option of plugin `fs` to set image save dir, like bellow:

`docker run -d -P -v rikka_files:/data --name rikka 7sdream/rikka -pwd 12345 -dir /data`

So you need't input a long mount path like `/go/src/github.com/7sDream/rikka/files`.

## Way 3: Use Docker Cloud Service Provider

For example, you can use free-plan of DaoCloud to deploy a Rikka server.

See [DaoCloud Deploy Guide][daocloud-guide] for detail.

## Use Other plugin

Main steps are the same.

See [Plugins Doc] for options for different plugins.

[version-zh]: https://github.com/7sDream/rikka/blob/master/deploy.zh.md

[daocloud-guide]: https://github.com/7sDream/rikka/wiki/%E5%9C%A8-DaoCloud-%E4%B8%8A%E5%85%8D%E8%B4%B9%E9%83%A8%E7%BD%B2-Rikka
[plugins-doc]: https://github.com/7sDream/rikka/tree/master/plugins
