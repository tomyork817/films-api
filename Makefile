build:
	docker-compose build vk-films-api

run: build
	docker-compose up vk-films-api

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5434/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5434/postgres?sslmode=disable' down

swag:
	swag init -g cmd/main.go

