install:
	docker run -it --rm -v $(PWD)/client:/app -w /app node:14.1-buster yarn install
up:
	docker-compose up -d
down:
	docker-compose down
log:
	docker-compose logs -f --tail=100
ps:
	docker-compose ps