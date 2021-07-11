CLIENT=client

install:
	docker build -t $(CLIENT) -f docker/nuxt/Dockerfile 
	docker run -it --rm -v $(PWD)/client:/app -w /app $(CLIENT) yarn install
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