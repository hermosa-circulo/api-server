default: run

run:
	go run main.go

start:
	docker-compose up -d

stop:
	docker-compose down -v

build:
	go build

docker: docker-build docker-push

docker-build:
	docker build -t hermosa/controller .

docker-push:
	docker push hermosa/controller
