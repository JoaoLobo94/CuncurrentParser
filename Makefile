include .env

install_golang_migrate:
	brew install golang-migrate

pullpostgres:
	docker pull postgres

runpostgres:
	docker run --name postgres_donut_db --env-file .env -d -p 5433:5432 postgres

createdb:
	docker exec -it postgres_donut_db createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) donut_db

dropdb:
	docker exec -it postgres_donut_db dropdb --username=$(POSTGRES_USER) donut_db

migratedb:
	migrate -path db/migration -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5433/donut_db?sslmode=disable" -verbose up

.PHONY: install_golang_migrate pull_postgres run_postgres createdb dropdb