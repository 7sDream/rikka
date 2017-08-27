FROM golang:latest

ARG VCS_REF
ARG VCS_URL
ARG BUILD_DATE
ARG VERSION

LABEL org.label-schema.schema-version="1.0" \
    org.label-schema.version=$VERSION \
    org.label-schema.build-date=$BUILD_DATE \
    org.label-schema.vcs-ref=$VCS_REF \
    org.label-schema.vcs-url=$VCS_URL \
    org.label-schema.vcs-type="Git" \
    org.label-schema.license="MIT" \
    org.label-schema.docker.dockerfile="/Dockerfile" \
    org.label-schema.name="Rikka"

MAINTAINER 7sDream "7seconddream@gmail.com"

WORKDIR $GOPATH/src/github.com/7sDream/rikka
ADD . $GOPATH/src/github.com/7sDream/rikka

RUN go get -v -d . && \
    go build -v . && \
    cp rikka $GOPATH/bin && \
    cp -R server $GOPATH/bin/ && \
    rm -rf $GOPATH/src

WORKDIR $GOPATH/bin

EXPOSE 80

ENTRYPOINT ["rikka"]
