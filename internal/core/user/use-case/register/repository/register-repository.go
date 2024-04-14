package repository

import (
	"context"

	"github.com/andersonjoseph/soundgo/internal/core/user/model"
)

type DtoSaveUser struct {
	Email    model.Email
	Username model.Username
	Password model.Password
}

type DtoUser struct {
	Id       int
	Email    string
	Username string
	Password string
}

type UserRepository interface {
	Save(ctx context.Context, dto DtoSaveUser) (DtoUser, error)
}
