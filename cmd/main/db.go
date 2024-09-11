package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func getPGURL() (string, error) {
	envVars := map[string]string{
		"DB_USER":     "",
		"DB_PASSWORD": "",
		"DB_HOST":     "",
		"DB_NAME":     "",
		"DB_PORT":     "",
	}

	for k := range envVars {
		v, ok := os.LookupEnv(k)

		if !ok {
			return "", fmt.Errorf("%s is missing from environment", k)
		}
		envVars[k] = v
	}

	//postgresql://user:password@host:port/name
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		envVars["DB_USER"],
		envVars["DB_PASSWORD"],
		envVars["DB_HOST"],
		envVars["DB_PORT"],
		envVars["DB_NAME"],
	), nil
}

func getDBConnection(ctx context.Context) (*pgxpool.Pool, error) {
	url, err := getPGURL()
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
