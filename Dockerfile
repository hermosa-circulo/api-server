FROM golang:1.8.3-alpine as builder
MAINTAINER bando-shintaro<bando142141@gmail.com>

ENV APP_DIR=/go/src/iga-controller

RUN apk update
RUN apk upgrade
RUN apk --no-cache add git musl-dev
RUN go get -u github.com/tools/godep

COPY . $APP_DIR
WORKDIR $APP_DIR
RUN godep get
RUN godep go build

FROM alpine:latest
COPY --from=builder /go/src/iga-controller/iga-controller /hermosa/iga-controller
CMD ["/hermosa/iga-controller"]
