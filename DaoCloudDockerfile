FROM golang:latest

MAINTAINER 7sDream "7seconddream@gmail.com"

WORKDIR $GOPATH/src/github.com/7sDream/rikka
ADD . $GOPATH/src/github.com/7sDream/rikka
RUN go build .

EXPOSE 80

ENTRYPOINT ["./rikka"]
