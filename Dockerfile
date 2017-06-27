FROM golang:1.8.3 as builder
MAINTAINER bando-shintaro<bando142141@gmail.com>

COPY . /go/src/api-server
WORKDIR /go/src/api-server
RUN go get -u github.com/tools/godep
RUN godep get
RUN godep go build

FROM alpine:latest
COPY --from=builder /go/src/api-server/api-server  .
CMD ["./api-server"]

