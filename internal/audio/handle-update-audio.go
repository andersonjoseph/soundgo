package audio

import (
	"context"
	"errors"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/shared"
)

func (h Handler) UpdateAudio(ctx context.Context, req *api.UpdateAudioInput, params api.UpdateAudioParams) (api.UpdateAudioRes, error) {
	a, err := h.repository.Get(ctx, params.ID)
	if errors.Is(err, shared.ErrNotFound) {
		return &api.UpdateAudioNotFound{}, nil
	}
	if err != nil {
		return nil, err
	}

	currUser, err := h.contextRequestHandler.GetUserID(ctx)
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
