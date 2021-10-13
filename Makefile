install:
	docker-compose build
	docker-compose run client yarn install
build-clinet:
	docker-compose run client yarn build
up:
	docker-compose -f docker-compose.yml up -d
up-prod:
	docker-compose -f docker-compose.prod.yml up -d
down:
	docker-compose down
log:
	docker-compose logs -f --tail=100
log-client:
	docker-compose logs -f --tail=100 client
ps:
	docker-compose ps
