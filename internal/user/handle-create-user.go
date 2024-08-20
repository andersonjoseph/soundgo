package user

import (
	"context"
	"errors"
	"log/slog"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/shared"
)

// POST /users
func (h Handler) CreateUser(ctx context.Context, req *api.UserInput) (api.CreateUserRes, error) {
	h.logger.Info("creating user",
		slog.Group(
			"input",
			"username", req.Username,
			"email", req.Email,
		),
	)

	userToSave, err := h.userInputToSaveInput(*req)
	if err != nil {
		return handleError(err)
	}

	savedUser, err := h.repository.Save(ctx, userToSave)
	if err != nil {
		return handleError(err)
	}

	return &api.User{
		ID:        savedUser.ID,
		Username:  savedUser.Username,
		Email:     savedUser.Email,
		CreatedAt: savedUser.CreatedAt,
	}, nil
}

func handleError(err error) (api.CreateUserRes, error) {
	switch {
	case errors.Is(err, shared.ErrAlreadyExists):
		return &api.CreateUserConflict{Error: ""}, nil

	default:
		return nil, err
	}
}

func (h Handler) userInputToSaveInput(ui api.UserInput) (SaveInput, error) {
	hashedPassword, err := h.hasher.Hash(ui.Password)
	if err != nil {
		return SaveInput{}, err
	}

	ID, err := shared.GenerateUUID()
	if err != nil {
		return SaveInput{}, err
	}

	return SaveInput{
		ID:       ID,
		Username: ui.Username,
		Email:    ui.Email,
		Password: hashedPassword,
	}, nil
}
