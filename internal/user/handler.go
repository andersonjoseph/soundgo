package user

import (
	"context"
	"log/slog"
	"time"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/shared"
)

type Entity struct {
	ID        string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Handler struct {
	repository Repository
	logger     *slog.Logger
	hasher     shared.PasswordHasher
}

func NewHandler(repository Repository, logger *slog.Logger, hasher shared.PasswordHasher) Handler {
	return Handler{
		repository: repository,
		logger:     logger,
		hasher:     hasher,
	}
}

// PATCH /users/{id}
func (h Handler) UpdateUser(ctx context.Context, req *api.UpdateUserInput, params api.UpdateUserParams) (api.UpdateUserRes, error) {
	h.logger.Info("updating user",
		slog.Group(
			"input",
			"username",
			req.Username,
		),
	)

	userToUpdate := UpdateInput{
		Username: req.GetUsername(),
	}
	updatedUser, err := h.repository.Update(ctx, params.ID, userToUpdate)

	if err != nil {
		h.logger.Error("error", "message", err.Error())
		return nil, err
	}

	return &api.User{
		ID:        updatedUser.ID,
		Username:  updatedUser.Username,
		Email:     updatedUser.Email,
		CreatedAt: updatedUser.CreatedAt,
	}, nil
}
