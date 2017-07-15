up:
	docker-compose up -d

db:
	docker-compose exec db mysql -uroot -pdocker iga

update:
	docker-compose up -d --build

log:
	docker-compose logs api
