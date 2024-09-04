package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/audio"
	"github.com/andersonjoseph/soundgo/internal/health"
	"github.com/andersonjoseph/soundgo/internal/password"
	"github.com/andersonjoseph/soundgo/internal/security"
	"github.com/andersonjoseph/soundgo/internal/session"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/andersonjoseph/soundgo/internal/user"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"
)

type SecurityHandler = security.Handler

type UserHandler = user.Handler
type SessionHandler = session.Handler
type PasswordHandler = password.Handler
type HealthHandler = health.Handler
type AudioHandler = audio.Handler

type serverHandler struct {
	UserHandler
	PasswordHandler
	SessionHandler
	HealthHandler
	AudioHandler
}

func hostInContextMiddleware(h http.Handler) (http.Handler, error) {
	host, ok := os.LookupEnv("HOST")
	if !ok {
		return nil, fmt.Errorf("HOST is missing from environment")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "host", host)

		h.ServeHTTP(w, r.WithContext(ctx))
	}), nil
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
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	pool, err := getDBConnection()
	if err != nil {
		logger.Error("error connecting to DB", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	hasher := shared.ScryptHasher{}
	userRepo := user.NewPGRepository(pool)
	sessionRepo := session.NewPGRepository(pool)
	JWTHandler := shared.JWTHandler{}
	requestContextHandler := shared.RequestContextHandler{}

	securityHandler := security.NewHandler(sessionRepo, JWTHandler, logger)
	audiosPath, ok := os.LookupEnv("AUDIOS_PATH")
	if !ok {
		logger.Error("AUDIOS_PATH is not present on environment", "error", err)
		os.Exit(1)
	}

	h := serverHandler{
		UserHandler: user.NewHandler(
			userRepo,
			logger,
			hasher,
			requestContextHandler,
		),
		SessionHandler: session.NewHandler(
			sessionRepo,
			userRepo,
			hasher,
			JWTHandler,
			logger,
			requestContextHandler,
		),
		PasswordHandler: password.NewHandler(
			password.NewPGRepository(pool),
			userRepo,
			hasher,
			shared.NewFakeEmailSender(logger),
			logger,
		),
		AudioHandler: audio.NewHandler(
			logger,
			audio.NewPGRepository(pool),
			audio.NewLocalFileRepository(audiosPath),
			requestContextHandler,
		),
	}

	srv, err := api.NewServer(h, securityHandler)
	if err != nil {
		logger.Error("error creating app server", "error", err)
		os.Exit(1)
	}

	handler := cors.AllowAll().Handler(srv)

	handler, err = hostInContextMiddleware(handler)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		logger.Error("PORT env variable missing from environment")
		os.Exit(1)
	}

	logger.Info("app started", "port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), handler); err != nil {
		logger.Error("error while starting http server", "error", err)
		os.Exit(1)
	}
}
