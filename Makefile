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
	atlas migrate apply \
		--dir "file://ent/migrate/migrations" \
		--url "postgres://$(DB_USERNAME):$(DB_PASSWORD)@localhost:15432/$(DB_NAME)?search_path=public&sslmode=disable"

.PHONY: db_migrate_diff
db_migrate_diff:
ifndef name
	@echo "Usage: make db_migrate_diff name=create_users_table"
	@echo ""
	@exit 1
endif
	atlas migrate diff $(name) \
		--dir "file://ent/migrate/migrations" \
		--to "ent://ent/schema" \
		--dev-url "docker://postgres/14/test?search_path=public"
