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
	GetByUser(ctx context.Context, userID string, after string, limit uint64, excludeHidden bool) ([]Entity, error)
	Delete(context.Context, string) error
	Update(ctx context.Context, audioID string, i UpdateInput) (Entity, error)
}
