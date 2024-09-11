package user

import (
	"context"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/shared"
)

// POST /users
func (h Handler) CreateUser(ctx context.Context, req *api.UserInput) (api.CreateUserRes, error) {
	userToSave, err := h.userInputToSaveInput(*req)
	if err != nil {
		return nil, err
	}

	savedUser, err := h.repository.Save(ctx, userToSave)
	if err != nil {
		return nil, err
	}

	return &api.User{
		ID:        savedUser.ID,
		Username:  savedUser.Username,
		Email:     savedUser.Email,
		CreatedAt: savedUser.CreatedAt,
	}, nil
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
