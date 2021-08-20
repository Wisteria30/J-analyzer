install:
	docker-compose build
	docker-compose run client yarn install
up:
	docker-compose up -d
down:
	docker-compose down
log:
	docker-compose logs -f --tail=100
log-client:
	docker-compose logs -f --tail=100 client
ps:
	docker-compose ps