package audio

import (
	"context"
	"errors"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/reqcontext"
	"github.com/andersonjoseph/soundgo/internal/shared"
)

func (h Handler) GetAudio(ctx context.Context, params api.GetAudioParams) (api.GetAudioRes, error) {
	e, err := h.repository.Get(ctx, params.ID)
	if errors.Is(err, shared.ErrNotFound) {
		return &api.GetAudioNotFound{}, nil
	}
	if err != nil {
		return nil, err
	}

	if e.Status == api.AudioStatusHidden {
		currentUserID, err := reqcontext.CurrentUserID.Get(ctx)

		if err != nil || currentUserID != e.UserID {
			return &api.GetAudioForbidden{}, nil
		}
	}

	return &api.Audio{
		ID:          e.ID,
		Title:       e.Title,
		Description: api.NewOptString(e.Description),
		CreatedAt:   e.CreatedAt,
		User:        e.UserID,
		Status:      e.Status,
		PlayCount:   e.Playcount,
	}, nil
}
