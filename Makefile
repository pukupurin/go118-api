include .env

.PHONY: install
install:
	curl -sSf https://atlasgo.sh | sh

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: run
run:
	docker exec -it go-ent-api ash -c "air"

db_migrate:
	docker compose run --rm go-ent-db-migration up

.PHONY: db_migrate_diff
db_migrate_diff:
ifndef name
	@echo "Usage: make db_migrate_diff name=create_users_table"
	@echo ""
	@exit 1
endif
	docker exec -it go-ent-api ash -c "go run -mod=mod ent/migrate/main.go $(name)"
