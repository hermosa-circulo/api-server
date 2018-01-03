default: proto run

run:
	go run cmd/iga-controller/main.go

start:
	docker-compose up -d

stop:
	docker-compose down -v

build: build-server build-client

build-server:
	go build -o iga-controller cmd/iga-controller/main.go

build-client:
	go build -o hctl cmd/hctl/main.go
	mv hctl ${GOPATH}/bin/

docker: docker-build docker-push

docker-build:
	docker build -t hermosa/controller .

docker-push:
	docker push hermosa/controller

proto:
	protoc --go_out=plugins=grpc:. api/iga.proto
