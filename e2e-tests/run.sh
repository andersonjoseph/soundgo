#!/bin/bash

echo "starting e2e services"
docker compose --profile test up -d --wait

echo "running migrations"
docker compose exec soundgo_env goose -dir ./migrations postgres ${DB} up

echo "starting http server"
docker compose exec -T -d soundgo_env bash -c "go run ./cmd/soundgo/main.go -db $DB -port $PORT &> /proc/1/fd/1"

echo "waiting for server to be ready..."
until docker compose exec -T soundgo_env curl -s -o /dev/null -w "%{http_code}" localhost:$PORT/api/v1/health | grep -q 200
do
  sleep 1
done

echo "running tests"
find . -name "*.hurl" -exec docker compose exec -T soundgo_env hurl --variable host=localhost:$PORT --test {} \;

echo "stopping http server"
for pid in $(docker compose exec -T soundgo_env ps -aux | awk '/soundgo_test/ {print $2}')
do
  docker compose exec -T -d soundgo_env kill -SIGTERM $pid
done

echo "stopping e2e services"
docker compose --profile test down -v
