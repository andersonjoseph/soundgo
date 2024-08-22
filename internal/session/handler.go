package session

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/andersonjoseph/soundgo/internal/user"
)

type Entity struct {
	ID    string
	Token string
}

type Handler struct {
	repository     Repository
	userRepository user.Repository
	logger         *slog.Logger
	hasher         shared.PasswordHasher
}

func NewHandler(repository Repository, userRepository user.Repository, hasher shared.PasswordHasher, logger *slog.Logger) Handler {
	return Handler{
		repository:     repository,
		userRepository: userRepository,
		logger:         logger,
		hasher:         hasher,
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

// DELETE /sessions
func (h Handler) DeleteSession(ctx context.Context) (api.DeleteSessionRes, error) {
	IDVal := ctx.Value("userID")

	ID, ok := IDVal.(string)
	if !ok {
		return nil, fmt.Errorf("bad ID")
	}

	h.logger.Info(
		"deleting session",
	)

	return &api.DeleteSessionNoContent{}, h.repository.Delete(ctx, ID)
}
