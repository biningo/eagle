FROM golang:1.16-alpine3.13 as builder
ENV GOPROXY=https://goproxy.io
WORKDIR /build
ADD . /build/
RUN CGO_ENABLED=0 go build -a -ldflags "-s -w" -o eagle /build/

FROM alpine:3.14.0
LABEL maintainer="icepan@aliyun.com"
ENV DOCKER_HOST="unix:///tmp/docker.sock"
EXPOSE 9999
COPY --from=builder /build/eagle /
ENTRYPOINT ["/eagle"]