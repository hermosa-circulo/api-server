db:
	docker-compose exec db mysql -uroot -pdocker iga
update:
	docker-compose up -d --build