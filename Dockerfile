FROM amd64/golang:latest as builder

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

MAINTAINER 7sDream "docker@7sdre.am"

ENV GO111MODULE=on

WORKDIR /go-modules
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v .

FROM amd64/alpine:latest

RUN mkdir -p /root/rikka/server/webserver

WORKDIR /root/rikka

COPY --from=builder /go-modules/rikka rikka
COPY --from=builder /go-modules/server/webserver/templates server/webserver/templates
COPY --from=builder /go-modules/server/webserver/static server/webserver/static

EXPOSE 80

ENTRYPOINT ["./rikka"]
