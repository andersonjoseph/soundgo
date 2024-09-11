package audio

import (
	"context"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/reqcontext"
)

func (h Handler) GetAudio(ctx context.Context, params api.GetAudioParams) (api.GetAudioRes, error) {
	e, err := h.repository.Get(ctx, params.ID)
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
