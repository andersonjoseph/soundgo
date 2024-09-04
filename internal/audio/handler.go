package audio

import (
	"log/slog"
	"time"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/shared"
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

type Handler struct {
	logger                *slog.Logger
	repository            Repository
	fileRepository        FileRepository
	contextRequestHandler shared.RequestContextHandler
}

func NewHandler(
	logger *slog.Logger,
	repo Repository,
	fileRepo FileRepository,
	ctxReqHandler shared.RequestContextHandler,
) Handler {
	return Handler{
		logger:                logger,
		repository:            repo,
		fileRepository:        fileRepo,
		contextRequestHandler: ctxReqHandler,
	}
}
