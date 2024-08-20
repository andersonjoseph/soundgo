package internaltest

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func getPGURL(t *testing.T) string {
	t.Helper()

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
			t.Fatalf("%s is missing from environment", k)
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
	)
}

func GenerateUUID(t *testing.T) string {
	t.Helper()
	IDHandler, err := uuid.NewV7()

	if err != nil {
		t.Fatal("error while generating ID handler: %w", err)
		return ""
	}

	ID, err := IDHandler.MarshalText()

	if err != nil {
		t.Fatal("error while generating ID handler: %w", err)
		return ""
	}

	return string(ID)
}

func GetPgPool(t *testing.T) *pgxpool.Pool {
	t.Helper()

	pool, err := pgxpool.New(context.TODO(), getPGURL(t))

	if err != nil {
		t.Fatal("error while connecting to the db: %w", err)
		return nil
	}

	if err = pool.Ping(context.TODO()); err != nil {
		t.Fatal("error while pinging to the db: %w", err)
		return nil
	}

	return pool
}
