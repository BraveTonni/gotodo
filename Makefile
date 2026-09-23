include .env
export

export PROJECT_ROOT=${shell cd}

env-up:
	@echo "Setting up the environment..."
	@docker-compose up -d todo-postgres
	@echo "Environment is up and running."

env-down:
	@echo "Tearing down the environment..."
	@docker-compose down todo-postgres
	@echo "Environment has been torn down."

migrate-create:
	@echo "Creating a new migration..."
	@docker-compose run --rm todo-postgres-migrate create -ext sql -dir /migrations -seq "${seq}"
	@echo "Migration created."

migrate-action:
	@docker-compose run --rm todo-postgres-migrate -path /migrations -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todo-postgres:5432/${POSTGRES_DB}?sslmode=disable "${action}"

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder