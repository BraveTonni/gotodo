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
