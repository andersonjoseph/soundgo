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
	repository            Repository
	logger                *slog.Logger
	hasher                shared.PasswordHasher
	contextRequestHandler shared.RequestContextHandler
}

func NewHandler(
	repository Repository,
	logger *slog.Logger,
	hasher shared.PasswordHasher,
	contextRequestHandler shared.RequestContextHandler,
) Handler {
	return Handler{
		repository:            repository,
		logger:                logger,
		hasher:                hasher,
		contextRequestHandler: contextRequestHandler,
	}
}
