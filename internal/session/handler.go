package session

import (
	"errors"
	"log/slog"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/andersonjoseph/soundgo/internal/user"
)

type Entity struct {
	ID     string
	Token  string
	UserID string
}

type Handler struct {
	repository            Repository
	userRepository        user.Repository
	logger                *slog.Logger
	hasher                shared.PasswordHasher
	jwtHandler            shared.JWTHandler
	contextRequestHandler shared.RequestContextHandler
}

func NewHandler(
	repository Repository,
	userRepository user.Repository,
	hasher shared.PasswordHasher,
	jwtHandler shared.JWTHandler,
	logger *slog.Logger,
	contextRequestHandler shared.RequestContextHandler,
) Handler {
	return Handler{
		repository:            repository,
		userRepository:        userRepository,
		logger:                logger,
		hasher:                hasher,
		jwtHandler:            jwtHandler,
		contextRequestHandler: contextRequestHandler,
	}
}

func (h Handler) handleError(err error) (api.CreateSessionRes, error) {
	switch {
	case errors.Is(err, shared.ErrUnauthorized):
		h.logger.Info(
			"unauthorized",
			"msg",
			err.Error(),
			slog.Group(
				"http_info",
				"status",
				401,
			),
		)
		return &api.Unauthorized{}, nil

	default:
		h.logger.Info(
			"internal server error",
			"msg",
			err.Error(),
			slog.Group(
				"http_info",
				"status",
				500,
			),
		)
		return nil, err
	}
}
