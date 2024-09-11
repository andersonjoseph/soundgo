package session

import (
	"context"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/reqcontext"
)

// DELETE /sessions
func (h Handler) DeleteSession(ctx context.Context) (api.DeleteSessionRes, error) {
	session, err := reqcontext.SessionID.Get(ctx)

	if err != nil {
		return nil, err
	}

	if err := h.repository.Delete(ctx, session); err != nil {
		return nil, err
	}

	return &api.DeleteSessionNoContent{}, nil
}
