FROM golang:1.7

MAINTAINER 7sDream "7seconddream@gmail.com"

WORKDIR $GOPATH/src/github.com/7sDream/rikka
ADD . $GOPATH/src/github.com/7sDream/rikka
RUN go build github.com/7sDream/rikka

EXPOSE 80

ENTRYPOINT $GOPATH/src/github.com/7sDream/rikka/rikka
CMD []
