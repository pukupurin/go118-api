.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: run
run:
	docker exec -it go-ent-api ash -c "air"
