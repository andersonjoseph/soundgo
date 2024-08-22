for file in $(find . -type f -name '*.hurl'); do
  docker compose --profile test up --wait -d
  docker compose exec soundgo_test goose -dir ./migrations up

  docker compose exec -d soundgo_test bash -c 'go run ./cmd/main/main.go >> /proc/1/fd/1 2>&1'

  sleep 3

  docker compose exec soundgo_test hurl --test --variable HOST=$HOST --variable PORT=$PORT $file

  docker compose logs soundgo_test
  docker compose --profile test down -v
done
