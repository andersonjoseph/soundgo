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
	session := ctx.Value("sessionID").(string)

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
