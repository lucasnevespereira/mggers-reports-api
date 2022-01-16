run:
	go run *.go

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

db-logs:
	docker-compose logs db

db-logs-error:
	docker-compose logs db | grep error
