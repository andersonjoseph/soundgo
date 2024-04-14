package repository

import (
	"context"

	"github.com/andersonjoseph/soundgo/internal/core/user/model"
)

type FindByUserDTO struct {
	ID       int
	Email    string
	Username string
}

type RequestResetPasswordRepository interface {
	FindByUsernameAndEmail(ctx context.Context, username model.Username, email model.Email) (FindByUserDTO, error)
}
