package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/jackc/pgx/v5/pgxpool"
)

type server struct {
	db                 *pgxpool.Pool
	logger             *slog.Logger
	handler            *http.ServeMux
	sessionTokenSecret []byte
}

func (s *server) Listen(addr string) error {
	return http.ListenAndServe(addr, s.handler)
}

func (s *server) handleError(ctx context.Context, err error, w http.ResponseWriter) {
	var httpCode int

	switch {
	case errors.Is(err, shared.ErrBadRequest), errors.Is(err, shared.ErrValidation):
		httpCode = http.StatusBadRequest

	case errors.Is(err, shared.ErrRecordAlreadyExists):
		httpCode = http.StatusConflict

	case errors.Is(err, shared.ErrUnauthorized):
		httpCode = http.StatusUnauthorized

	case errors.Is(err, shared.ErrNotFound):
		httpCode = http.StatusNotFound

	default:
		httpCode = http.StatusInternalServerError
	}

	s.logger.Debug("", "http_code:", httpCode)

	s.logger.LogAttrs(
		ctx,
		slog.LevelError,
		err.Error(),
		slog.Group(
			"req",
			"id", ctx.Value("req_id"),
			"path", ctx.Value("req_path"),
		),
		slog.Int("http_code", httpCode),
	)

	if httpCode == http.StatusInternalServerError {
		jsonError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	jsonError(w, err.Error(), httpCode)
}

func NewServer(
	db *pgxpool.Pool,
	logger *slog.Logger,
	sessionTokenSecret []byte,
) server {
	s := server{
		db:                 db,
		logger:             logger,
		handler:            http.NewServeMux(),
		sessionTokenSecret: sessionTokenSecret,
	}

	s.registerRoutes()

	return s
}
