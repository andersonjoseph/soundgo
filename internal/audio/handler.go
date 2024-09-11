package audio

import (
	"context"
	"log/slog"
	"time"

	"github.com/andersonjoseph/soundgo/internal/api"
)

type Entity struct {
	ID          string
	Title       string
	Description string
	Playcount   int
	Status      api.AudioStatus
	UserID      string
	CreatedAt   time.Time
}

type PlayCountHandler interface {
	Add(ctx context.Context, userID string, audio Entity) error
}

type Handler struct {
	logger           *slog.Logger
	repository       Repository
	fileRepository   FileRepository
	playCountHandler PlayCountHandler
}

func NewHandler(
	logger *slog.Logger,
	repo Repository,
	fileRepo FileRepository,
	playCountHandler PlayCountHandler,
) Handler {
	return Handler{
		logger:           logger,
		repository:       repo,
		fileRepository:   fileRepo,
		playCountHandler: playCountHandler,
	}
}
