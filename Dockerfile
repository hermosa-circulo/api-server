FROM golang:1.9.2-alpine as builder

ENV APP_DIR=/go/src/github.com/hermosa-circulo/iga-controller

RUN apk update
RUN apk upgrade
RUN apk --no-cache add git musl-dev make

COPY . $APP_DIR
WORKDIR $APP_DIR
RUN make build-server

FROM alpine:latest
COPY --from=builder /go/src/github.com/hermosa-circulo/iga-controller/iga-controller /hermosa/iga-controller
CMD ["/hermosa/iga-controller"]
