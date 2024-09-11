package user

import (
	"context"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/reqcontext"
)

// PATCH /users/{id}
func (h Handler) UpdateUser(ctx context.Context, req *api.UpdateUserInput, params api.UpdateUserParams) (api.UpdateUserRes, error) {
	currentUserID, err := reqcontext.CurrentUserID.Get(ctx)
	if err != nil {
		return nil, err
	}

	if currentUserID != params.ID {
		return &api.Unauthorized{}, nil
	}

	updatedUser, err := h.repository.Update(ctx, params.ID, UpdateInput{
		Username: req.GetUsername(),
	})

	if err != nil {
		return nil, err
	}

	return &api.User{
		ID:        updatedUser.ID,
		Username:  updatedUser.Username,
		Email:     updatedUser.Email,
		CreatedAt: updatedUser.CreatedAt,
	}, nil
}
