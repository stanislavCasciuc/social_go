include .env
MIGRATION_PATH = ./cmd/migrate/migrations

build:
	@go build -o bin/social cmd/api/*.go
run: build
	@./bin/social
migration:
	@migrate create -seq -ext sql -dir $(MIGRATION_PATH) $(filter-out $@,$(MAKECMDGOALS))
migrate-up:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) up	
migrate-down:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))	$(filter-out $@,$(MAKECMDGOALS))
seed:
	@go run cmd/migrate/seed/main.go
