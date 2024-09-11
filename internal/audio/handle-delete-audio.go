package audio

import (
	"context"
	"errors"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/reqcontext"
	"github.com/andersonjoseph/soundgo/internal/shared"
)

func (h Handler) DeleteAudio(ctx context.Context, params api.DeleteAudioParams) (api.DeleteAudioRes, error) {
	a, err := h.repository.Get(ctx, params.ID)
	if errors.Is(err, shared.ErrNotFound) {
		return &api.DeleteAudioNotFound{}, nil
	}
	if err != nil {
		return nil, err
	}

	currUser, err := reqcontext.CurrentUserID.Get(ctx)
	if err != nil {
		return nil, err
	}
	if currUser != a.UserID {
		return &api.DeleteAudioForbidden{}, nil
	}

	err = h.repository.Delete(ctx, params.ID)
	if errors.Is(err, shared.ErrNotFound) {
		return &api.DeleteAudioNotFound{}, nil
	}
	if err != nil {
		return nil, err
	}

	err = h.fileRepository.Remove(ctx, params.ID)
	if errors.Is(err, shared.ErrNotFound) {
		return &api.DeleteAudioNotFound{}, nil
	}
	if err != nil {
		return nil, err
	}

	return &api.DeleteAudioNoContent{}, nil
}
