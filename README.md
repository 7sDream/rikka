# Rikka - A simple photo share website.

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

But, if you close this page, you have no way to find it back except from browser history.

## Deploy

### Method 1: Build your self on your VPS

1. `go get github.com/7sDream/rikka`
2. `cd $GOPATH/src/github.com/7sDream/rikka`
3. `go build github.com/7sDream/rikka`
4. `./rikka --port 80 --pwd yourpassword`

Then you can view your website and use the password you set to upload and share photo.

### Method 2: Use docker

working on it.
