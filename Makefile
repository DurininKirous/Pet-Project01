build:
	go build -o app ./cmd/main.go

test:
	go test ./... -v

run:
	go run ./cmd/main.go

docker-build:
	docker-compose build

up:
	docker-compose up

up-detached:
	docker-compose up -d

down:
	docker-compose down

clean:
	docker-compose down -v

logs:
	docker-compose logs

rebuild:
	docker-compose down -v
	docker-compose up --build

.PHONY: build test run docker-build up up-detached down clean logs rebuild

