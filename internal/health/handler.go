package health

import (
	"context"
	"log/slog"

	"github.com/andersonjoseph/soundgo/internal/api"
)

type Handler struct {
	logger *slog.Logger
}

func (h Handler) CheckHealth(ctx context.Context) (api.CheckHealthRes, error) {
	return &api.CheckHealthOK{
		Status: api.CheckHealthOKStatusHealthy,
	}, nil
}
