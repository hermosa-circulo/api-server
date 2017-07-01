FROM golang:1.8.3-alpine as builder
MAINTAINER bando-shintaro<bando142141@gmail.com>

RUN apk update
RUN apk upgrade
RUN apk --no-cache add git musl-dev
RUN go get -u github.com/tools/godep
COPY . /go/src/api-server
WORKDIR /go/src/api-server

RUN godep get
RUN godep go build

FROM alpine:latest
COPY --from=builder /go/src/api-server/api-server  /hermosa/api-server
CMD ["/hermosa/api-server"]
