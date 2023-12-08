.PHONY = all up run migrate-up test

MIGRATE = $(HOME)/go/bin/migrate -database "postgres://postgres:pgpassword@localhost:5432/rinha?sslmode=disable" -path ./database/migrations

all: build

build:
	@go build -o bin/rinha cmd/api/main.go

start-up:
	@docker-compose up --remove-orphans

dev:
	@go run cmd/web/main.go

migrate:
	@$(HOME)/go/bin/migrate create -ext sql -dir database/migrations $(filter-out $@,$(MAKECMDGOALS))

up:
	@go run cmd/migrate/main.go up

down:
	@go run cmd/migrate/main.go down
