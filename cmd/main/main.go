package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/audio"
	"github.com/andersonjoseph/soundgo/internal/health"
	"github.com/andersonjoseph/soundgo/internal/password"
	"github.com/andersonjoseph/soundgo/internal/security"
	"github.com/andersonjoseph/soundgo/internal/session"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/andersonjoseph/soundgo/internal/user"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ogen-go/ogen/ogenerrors"
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

func createErrorHandler(log *slog.Logger) ogenerrors.ErrorHandler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
		var code = http.StatusInternalServerError
		var ogenErr ogenerrors.Error

		switch {
		case errors.Is(err, shared.ErrAlreadyExists):
			code = http.StatusConflict

		case errors.Is(err, shared.ErrNotFound):
			code = http.StatusNotFound

		case errors.Is(err, shared.ErrBadInput):
			code = http.StatusBadRequest

		case errors.As(err, &ogenErr):
			code = ogenErr.Code()
		}
		msg := err.Error()
		log.Info(http.StatusText(code), "msg", msg)

		if code == http.StatusInternalServerError {
			log.Error("internal server error", "msg", msg)
			msg = "Internal server error"
		}

		w.WriteHeader(code)
		_, _ = io.WriteString(w, fmt.Sprintf(`{"error": "%v"}`, msg))
	}
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
	audioRepo := audio.NewPGRepository(pool)
	sessionRepo := session.NewPGRepository(pool)

	JWTHandler := shared.JWTHandler{}
	securityHandler := security.NewHandler(sessionRepo, JWTHandler, logger)

	if err != nil {
		logger.Error("error creating big cache play count handler", "error", err)
		os.Exit(1)
	}

	playCountSaveIntervalStr, ok := os.LookupEnv("PLAY_COUNT_SAVE_INTERVAL")
	if !ok {
		logger.Error("PLAY_COUNT_SAVE_INTERVAL is not present on environment", "error", err)
		os.Exit(1)
	}

	playCountSaveInterval, err := strconv.ParseInt(playCountSaveIntervalStr, 10, 64)
	if err != nil {
		logger.Error("PLAY_COUNT_SAVE_INTERVAL is not a valid integer", "error", err)
		os.Exit(1)
	}

	playCountHandler := audio.NewMemoryPlayCountHandler(context.Background(), 1<<17, audioRepo, time.Second*time.Duration(playCountSaveInterval), logger)

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
		),
		SessionHandler: session.NewHandler(
			sessionRepo,
			userRepo,
			hasher,
			JWTHandler,
		),
		PasswordHandler: password.NewHandler(
			password.NewPGRepository(pool),
			userRepo,
			hasher,
			shared.NewFakeEmailSender(logger),
		),
		AudioHandler: audio.NewHandler(
			audioRepo,
			audio.NewLocalFileRepository(audiosPath),
			playCountHandler,
		),
	}

	srv, err := api.NewServer(h, securityHandler, api.WithErrorHandler(createErrorHandler(logger)))

	if err != nil {
		logger.Error("error creating app server", "error", err)
		os.Exit(1)
	}

	var handler http.Handler = srv
	handler = LogRequestMiddlware(srv, srv, logger)

	handler = clientFingerprintMiddleware(handler, srv)
	handler = cors.AllowAll().Handler(handler)

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
