docker compose exec soundgo_test goose -dir ./migrations up;

for file in $(find . -type f -name '*.hurl'); do
	docker compose --profile test up --wait -d
    docker compose exec soundgo_test goose -dir ./migrations up;

    docker compose exec soundgo_test hurl --test --variable HOST=$HOST --variable PORT=$PORT $file

    docker compose --profile test down -v
done
