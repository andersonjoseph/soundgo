package audio

import (
	"context"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/reqcontext"
)

func (h Handler) UpdateAudio(ctx context.Context, req *api.UpdateAudioInput, params api.UpdateAudioParams) (api.UpdateAudioRes, error) {
	a, err := h.repository.Get(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	currUser, err := reqcontext.CurrentUserID.Get(ctx)
	if err != nil {
		return nil, err
	}
	if currUser != a.UserID {
		return &api.UpdateAudioForbidden{}, nil
	}

	a, err = h.repository.Update(ctx, params.ID, UpdateInput{
		Title:       req.Title.Value,
		Description: req.Description.Value,
		Status:      req.Status.Value,
	})

	if err != nil {
		return nil, err
	}

	return &api.Audio{
		ID:          a.ID,
		Title:       a.Title,
		Description: api.NewOptString(a.Description),
		CreatedAt:   a.CreatedAt,
		PlayCount:   a.Playcount,
		User:        a.UserID,
		Status:      a.Status,
	}, nil
}
