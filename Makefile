include .env
export

CONTAINER_NAME=postgres_db
IMAGE_NAME=my_postgres
APP_NAME=myapp

.PHONY: deps build run docker-build docker-run docker-stop docker-logs psql clean setup

deps:
	go mod tidy

build:
	go build -o $(APP_NAME) main.go

run:
	go run main.go

docker-build:
	docker build -t $(IMAGE_NAME) .

docker-run: docker-build
	docker run -d \
		--name $(CONTAINER_NAME) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-p $(DB_PORT):5432 \
		$(IMAGE_NAME)

docker-stop:
	docker stop $(CONTAINER_NAME)

docker-logs:
	docker logs -f $(CONTAINER_NAME)

psql:
	docker exec -it $(CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME)

docker-clean:
	docker rm -f $(CONTAINER_NAME) || true
	docker rmi -f $(IMAGE_NAME) || true

clean: docker-clean
	rm -f $(APP_NAME)

setup: docker-run
	@echo "Waiting for database..."
	@sleep 3
	@echo "Ready! Run: make run"