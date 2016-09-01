FROM golang:1.6

MAINTAINER 7sDream "7seconddream@gmail.com"

WORKDIR $GOPATH/src/rikka
ADD . $GOPATH/src/rikka
RUN go build rikka

EXPOSE 80

CMD $GOPATH/src/rikka/rikka
