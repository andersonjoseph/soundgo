PORT := 3000
DB_CONN_STRING := postgres://admin:password@db:5432/soundgo

TEST_DB_CONN_STRING := postgres://admin:password@test_db:5432/soundgo_test
TEST_PORT := 3001

start-env:
	docker compose --profile dev up -d --wait

stop-env:
	docker compose --profile dev down -v

shell:
	docker compose exec soundgo_env bash

debug:
	docker compose --profile test up -d --wait
	make run-test-migrations
	- docker compose exec soundgo_env bash -c "dlv debug ./cmd/soundgo/main.go -- -db ${TEST_DB_CONN_STRING} -port ${TEST_PORT} &> /proc/1/fd/1"
	docker compose --profile test down -v
	docker compose restart soundgo_env

t:
	docker compose exec soundgo_env go test --short ./...

tests:
	docker compose --profile test up -d --wait
	make run-test-migrations
	- docker compose exec soundgo_env bash -c "DB=${TEST_DB_CONN_STRING} go test -count=1 ./..."
	docker compose --profile test down -v

e2e:
	DB=${TEST_DB_CONN_STRING} PORT=${TEST_PORT} bash ./e2e-tests/run.sh

T: 
	make t && make tests && make e2e

fmt:
	docker compose exec soundgo_env go fmt ./...

server:
	docker compose exec soundgo_env bash -c "air -- -db ${DB_CONN_STRING} -port ${PORT}"

preview-docs:
	docker compose exec soundgo_env redocly preview-docs ./docs/spec.yaml --host 0.0.0.0

run-migrations:
	docker compose exec soundgo_env goose -dir ./migrations postgres ${DB_CONN_STRING} up

run-test-migrations:
	docker compose exec soundgo_env goose -dir ./migrations postgres ${TEST_DB_CONN_STRING} up

generate-migrations:
	docker compose exec soundgo_env goose -s -dir ./migrations postgres ${DB_CONN_STRING} create $(name) sql
