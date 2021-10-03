.SILENT:


build-server: 
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/weather-app/main.go

stop-todo-db:
	docker stop todo-rest-api_db_1

run-db:
	docker run --name=weather-db -e POSTGRES_PASSWORD=klaus -p 5436:5432 -d --rm postgres

migrate-down:
	migrate -path ./schema -database postgres://postgres:klaus@localhost:5436/postgres?sslmode=disable down

exec-db:
	docker exec -it weather-db  /bin/bash

run-server: 
	go run cmd/weather-app/main.go