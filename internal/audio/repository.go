package audio

import (
	"context"

	"github.com/andersonjoseph/soundgo/internal/api"
)

type SaveInput struct {
	ID          string
	Title       string
	Description string
	UserID      string
	Status      api.AudioStatus
}

type Repository interface {
	Save(ctx context.Context, i SaveInput) (Entity, error)
	Get(context.Context, string) (Entity, error)
}
