package repository

import (
	"context"

	"github.com/andersonjoseph/soundgo/internal/core/user/model"
)

type ResetPasswordRepository interface {
	FindPasswordByID(ctx context.Context, id int) (string, error)
	UpdatePassword(ctx context.Context, id int, password model.Password) error
}
