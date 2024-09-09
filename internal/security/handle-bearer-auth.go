package security

import (
	"context"
	"errors"
	"log/slog"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/reqcontext"
	"github.com/andersonjoseph/soundgo/internal/session"
	"github.com/andersonjoseph/soundgo/internal/shared"
)

type Handler struct {
	logger            *slog.Logger
	sessionRepository session.Repository
	jwtHandler        shared.JWTHandler
}

func NewHandler(sessionRepository session.Repository, jwtHandler shared.JWTHandler, logger *slog.Logger) Handler {
	return Handler{
		logger:            logger,
		sessionRepository: sessionRepository,
	}
}

func (h Handler) HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error) {
	token, err := h.jwtHandler.ParseToken(t.Token)

	if err != nil {
		h.logger.Error(
			"error parsing JWT",
			"msg",
			err.Error(),
			slog.Group("http_info", "status", 401),
		)
		return nil, err
	}

	session, err := h.sessionRepository.GetByID(ctx, token.Subject())
	if errors.Is(err, shared.ErrNotFound) {
		h.logger.Info(
			"session does not exist",
			"msg",
			err.Error(),
			slog.Group("http_info", "status", 401),
		)
		return nil, err
	}
	if err != nil {
		h.logger.Error(
			"error while getting session",
			"msg",
			err.Error(),
			slog.Group("http_info", "status", 401),
		)
		return nil, err
	}

	ctxHandler := reqcontext.Handler{}

	ctx = ctxHandler.SetUserID(ctx, session.UserID)
	ctx = ctxHandler.SetSessionID(ctx, session.ID)
	return ctx, nil
}
