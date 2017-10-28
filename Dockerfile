FROM golang:1.9.2-alpine as builder
MAINTAINER bando-shintaro<bando142141@gmail.com>

ENV APP_DIR=/go/src/github.com/hermosa-circulo/iga-controller

RUN apk update
RUN apk upgrade
RUN apk --no-cache add git musl-dev

COPY . $APP_DIR
WORKDIR $APP_DIR
RUN go build

FROM alpine:latest
COPY --from=builder /go/src/github.com/hermosa-circulo/iga-controller/iga-controller /hermosa/iga-controller
CMD ["/hermosa/iga-controller"]
