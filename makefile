# Makefile for Cassandra-based Job Scheduler

# Variables
DOCKER_COMPOSE = docker-compose
CASSANDRA_CONTAINER = cassandra_node
MIGRATION_FILE = internal/migrations/init.cql

# Default target
all: up migrate

dev:
	go run cmd/server/main.go

# Start the Cassandra container
up:
	$(DOCKER_COMPOSE) up -d

# Stop the Cassandra container
down:
	$(DOCKER_COMPOSE) down

# Stop the Cassandra container and remove volumes
clean:
	$(DOCKER_COMPOSE) down -v

# Apply migrations
migrate:
	go run cmd/migrate/main.go

cassandra-init:
	@echo "Initializing Cassandra..."
	@./cassandra-init.sh

# Show Cassandra logs
logs:
	$(DOCKER_COMPOSE) logs -f cassandra

# Enter Cassandra CQL shell
cqlsh:
	docker exec -it $(CASSANDRA_CONTAINER) cqlsh

# Rebuild and restart the container
rebuild: down up

# Help target
help:
	@echo "Available targets:"
	@echo "  up        - Start the Cassandra container"
	@echo "  down      - Stop the Cassandra container"
	@echo "  clean     - Stop the container and remove volumes"
	@echo "  migrate   - Apply database migrations"
	@echo "  logs      - Show Cassandra logs"
	@echo "  cqlsh     - Enter Cassandra CQL shell"
	@echo "  rebuild   - Rebuild and restart the container"
	@echo "  all       - Start container and apply migrations (default)"

.PHONY: all up down clean migrate logs cqlsh rebuild help