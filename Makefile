install:
	docker-compose run client yarn install
build-client:
	docker-compose run client yarn build
migration:
	docker-compose up -d db
	docker-compose run server go run tools/migrate.go
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
