package audio

import (
	"log/slog"
	"time"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/reqcontext"
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
	contextRequestHandler reqcontext.Handler
}

func NewHandler(
	logger *slog.Logger,
	repo Repository,
	fileRepo FileRepository,
	ctxReqHandler reqcontext.Handler,
) Handler {
	return Handler{
		logger:                logger,
		repository:            repo,
		fileRepository:        fileRepo,
		contextRequestHandler: ctxReqHandler,
	}
}
