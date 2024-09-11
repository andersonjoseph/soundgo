package user

import (
	"context"
	"errors"
	"log/slog"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/reqcontext"
	"github.com/andersonjoseph/soundgo/internal/shared"
)

// PATCH /users/{id}
func (h Handler) UpdateUser(ctx context.Context, req *api.UpdateUserInput, params api.UpdateUserParams) (api.UpdateUserRes, error) {
	h.logger.Info("updating user",
		slog.Group(
			"input",
			"username",
			req.Username,
		),
	)

	currentUserID, err := reqcontext.CurrentUserID.Get(ctx)
	if err != nil {
		return nil, err
	}

	if currentUserID != params.ID {
		h.logger.Info("user attempted to update other user",
			"user",
			currentUserID,
			slog.Group(
				"http_info",
				"status",
				401,
			),
		)
		return &api.Unauthorized{}, nil
	}

	updatedUser, err := h.repository.Update(ctx, params.ID, UpdateInput{
		Username: req.GetUsername(),
	},
	)

	if errors.Is(err, shared.ErrNotFound) {
		h.logger.Info("user not found",
			"user",
			currentUserID,
			slog.Group(
				"http_info",
				"status",
				404,
			),
		)
		return &api.NotFound{}, nil
	}

	if err != nil {
		h.logger.Error("error", "message", err.Error())
		return nil, err
	}

	h.logger.Info("user updated",
		slog.Group(
			"http_info",
			"status",
			200,
		),
	)

	return &api.User{
		ID:        updatedUser.ID,
		Username:  updatedUser.Username,
		Email:     updatedUser.Email,
		CreatedAt: updatedUser.CreatedAt,
	}, nil
}
