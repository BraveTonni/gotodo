include .env
export

export PROJECT_ROOT=${shell cd}

env-up:
	@echo "Setting up the environment..."
	@docker-compose up -d todo-postgres
	@docker-compose up -d port-forwarder
	@echo "Environment is up and running."

env-down:
	@echo "Tearing down the environment..."
	@docker-compose down todo-postgres
	@docker-compose down port-forwarder
	@echo "Environment has been torn down."

migrate-create:
	@echo "Creating a new migration..."
	@docker-compose run --rm todo-postgres-migrate create -ext sql -dir /migrations -seq "${seq}"
	@echo "Migration created."

migrate-action:
	@docker-compose run --rm todo-postgres-migrate -path /migrations -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todo-postgres:5432/${POSTGRES_DB}?sslmode=disable "${action}"

app-run:
	go mod tidy && go run cmd/todo/main.go