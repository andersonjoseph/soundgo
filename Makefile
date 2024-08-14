start-env:
	docker compose --profile dev up -d --wait

stop-env:
	docker compose --profile dev down -v

shell:
	docker compose exec soundgo_env bash

preview-docs:
	docker compose exec soundgo_env bash -c 'redocly lint ./openapi/spec.yaml && redocly preview-docs -h 0.0.0.0 ./openapi/spec.yaml'