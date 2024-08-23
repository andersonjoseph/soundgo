ifneq (,$(wildcard ./.env))
    include .env
    export
endif

build-env:
	docker compose --profile dev build

start-env:
	docker compose --profile dev up -d --wait

stop-env:
	docker compose --profile dev down -v

shell:
	docker compose exec soundgo_env bash

preview-docs:
	docker compose exec soundgo_env bash -c 'redocly lint ./openapi/spec.yaml && redocly preview-docs -h 0.0.0.0 ./openapi/spec.yaml'

generate-openapi:
	docker compose exec soundgo_env bash -c 'redocly lint ./openapi/spec.yaml && redocly bundle -o ./openapi/spec.bundle.yaml ./openapi/spec.yaml && go generate ./open-api-gen.go'

fmt:
	docker compose exec soundgo_env go fmt ./... 

create-migration:
	@if [ -z "$(name)" ]; then \
			echo "Error: Please provide a migration name. Usage: make create-migration name=<migration_name>"; \
			exit 1; \
	fi; \
	docker compose exec soundgo_env goose -dir ./migrations create $(name) sql

test:
	docker compose --profile test up --wait -d
	@trap 'docker compose --profile test down -v' EXIT; \
	docker compose exec soundgo_test goose -dir ./migrations up; \
	docker compose exec soundgo_test go test ./...

e2e-tests:
	HOST=${TEST_HOST} PORT=${TEST_PORT} bash ./e2e/run.sh
		
debug:
	docker compose --profile test up --wait -d
	@trap 'docker compose --profile test down -v' EXIT; \
	docker compose exec soundgo_test goose -dir ./migrations up; \
	docker compose exec -it soundgo_test bash -c "dlv debug ./cmd/main/main.go"

request:
	@if [ -z "$(file)" ]; then \
			echo "Error: Please provide a migration name. Usage: make request file=<hurl_file>"; \
			exit 1; \
	fi; \
	docker compose exec soundgo_test goose -dir ./migrations up
	trap 'docker compose down -v test_db && docker compose up --wait -d test_db' EXIT; \
	docker compose exec soundgo_test hurl --test --variable HOST=${TEST_HOST} --variable PORT=${TEST_PORT} $(file)
