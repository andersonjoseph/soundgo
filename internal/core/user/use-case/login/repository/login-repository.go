package repository

import (
	"context"

	"github.com/andersonjoseph/soundgo/internal/core/user/model"
)

type DtoPassword struct {
	UserID int
	Value  string
}

type DtoCreateSession struct {
	UserID       int
	CreationDate string
}

type DtoSession struct {
	ID           int
	CreationDate string
}

type LoginRepository interface {
	FindPasswordByUsername(ctx context.Context, username model.Username) (DtoPassword, error)
	SaveSession(ctx context.Context, dto DtoCreateSession) (DtoSession, error)
}
