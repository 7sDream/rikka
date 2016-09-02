FROM golang:1.6

MAINTAINER 7sDream "7seconddream@gmail.com"

WORKDIR $GOPATH/src/github.com/7sDream/rikka
ADD . $GOPATH/src/github.com/7sDream/rikka
RUN go build github.com/7sDream/rikka

EXPOSE 80

CMD $GOPATH/src/github.com/7sDream/rikka/rikka
