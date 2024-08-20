package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/health"
	"github.com/andersonjoseph/soundgo/internal/password"
	"github.com/andersonjoseph/soundgo/internal/session"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/andersonjoseph/soundgo/internal/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserHandler = user.Handler
type SessionHandler = session.Handler
type PasswordHandler = password.Handler
type HealthHandler = health.Handler

type serverHandler struct {
	UserHandler
	PasswordHandler
	SessionHandler
	HealthHandler
}

type securityHandler struct{}

func (h securityHandler) HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error) {
	return context.TODO(), nil
}

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

func getDBConnection() (*pgxpool.Pool, error) {
	url, err := getPGURL()
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func main() {
	logger := slog.New(slog.Default().Handler())

	pool, err := getDBConnection()
	if err != nil {
		logger.Error("error connecting to DB", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	h := serverHandler{
		UserHandler: user.NewHandler(
			user.NewPGRepository(pool),
			logger,
			shared.ScryptHasher{},
		),
	}

	srv, err := api.NewServer(h, securityHandler{})
	if err != nil {
		logger.Error("error creating app server", "error", err)
		os.Exit(1)
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		logger.Error("PORT env variable missing from environment")
		os.Exit(1)
	}

	logger.Info("app started", "port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), srv); err != nil {
		logger.Error("error while starting http server", "error", err)
		os.Exit(1)
	}
}
