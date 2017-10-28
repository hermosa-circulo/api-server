db:
	docker-compose exec db mysql -uroot -pdocker iga
start:
	docker-compose up -d

run:
	go run main.go

stop:
	docker-compose down -v

build:
	go build


docker: docker-build docker-push

docker-build:
	docker build -t hermosa/controller .

docker-push:
	docker push hermosa/controller
