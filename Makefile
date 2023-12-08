all: build

build:
	@go build -o bin/rinha cmd/api/main.go

docker-up:
	@docker-compose up --remove-orphans

run:
	@go run cmd/web/main.go

migrate:
	@$(HOME)/go/bin/migrate create -ext sql -dir internal/database/migrations $(filter-out $@,$(MAKECMDGOALS))

up:
	@go run cmd/migrate/main.go up

down:
	@go run cmd/migrate/main.go down

watch:
	@if command -v air > /dev/null; then \
			air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
					air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi


.PHONY: all build start-up dev migrate up down
