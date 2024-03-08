-include .env

build:
	docker compose build

up:
	docker compose up -d

log:
	docker compose logs

db:
	docker exec -it $(POSTGRES_CONTAINER_HOST) psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)

go:
	docker exec -it $(GO_CONTAINER_HOST) /bin/sh