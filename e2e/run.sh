docker compose --profile test up --wait -d
docker compose exec -d soundgo_test bash -c 'go run ./cmd/main/*.go >> /proc/1/fd/1 2>&1'
sleep 3

if [ -z ${FILES} ]
then
    FILES=$(find . -type f -name '*.hurl')
fi

for file in $FILES; do
  docker compose up test_db --wait -d
  docker compose exec soundgo_test goose -dir ./migrations up

  trap "docker compose logs soundgo_test && exit 1" ERR 
  docker compose exec soundgo_test hurl --test --variable HOST=$HOST --variable PORT=$PORT $file
  
  docker compose down -v test_db
done
