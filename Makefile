run:
	go run cmd/main.go

dev:
	docker-compose up

start:
	docker-compose up -d

stop:
	docker-compose down

rebuild:
	docker-compose up --build

api-logs:
	docker-compose logs api

api-logs-error:
	docker-compose logs api | grep error

api-logs-info:
	docker-compose logs api | grep info

