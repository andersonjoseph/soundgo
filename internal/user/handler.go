package user

import (
	"log/slog"
	"time"

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
	hasher     shared.PasswordHasher
}

func NewHandler(
	repository Repository,
	logger *slog.Logger,
	hasher shared.PasswordHasher,
) Handler {
	return Handler{
		repository: repository,
		hasher:     hasher,
	}
}
