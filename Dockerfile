# ---------------
# Build Stage
# ---------------

FROM amd64/golang:1 as builder

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
    org.label-schema.name="Rikka" \
    maintainer="docker@7sdre.am"

ENV GO111MODULE=on

WORKDIR /src
COPY . .

RUN go env -w GOPROXY="https://goproxy.io,direct" && \
    go env -w GOSUMDB="gosum.io+ce6e7565+AY5qEHUk/qmHc5btzW45JVoENfazw8LielDsaI+lEbq6" && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v .

# ---------------
# Final Stage
# ---------------

FROM amd64/alpine:3

WORKDIR /root/rikka

COPY --from=builder /src/rikka rikka
COPY --from=builder /src/server/webserver/templates server/webserver/templates
COPY --from=builder /src/server/webserver/static server/webserver/static

EXPOSE 80

ENTRYPOINT ["./rikka"]
