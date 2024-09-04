package session

import (
	"context"
	"log/slog"

	"github.com/andersonjoseph/soundgo/internal/api"
)

// DELETE /sessions
func (h Handler) DeleteSession(ctx context.Context) (api.DeleteSessionRes, error) {
	h.logger.Info(
		"deleting session",
	)
	session, err := h.contextRequestHandler.GetSessionID(ctx)

	if err != nil {
		return nil, err
	}

	if err := h.repository.Delete(ctx, session); err != nil {
		h.logger.Error(
			"error deleting session",
			"msg",
			err.Error(),
			slog.Group("http_info", "status", 500),
		)
		return nil, err
	}

	return &api.DeleteSessionNoContent{}, nil
}
