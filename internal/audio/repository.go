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
	Status      api.AudioInputMultipartStatus
}

type UpdateInput struct {
	Title       string
	Description string
	Status      api.UpdateAudioInputStatus
	PlayCount   uint64
}

type Repository interface {
	Save(ctx context.Context, i SaveInput) (Entity, error)
	Get(context.Context, string) (Entity, error)
	Delete(context.Context, string) error
	Update(ctx context.Context, userID string, i UpdateInput) (Entity, error)
}
