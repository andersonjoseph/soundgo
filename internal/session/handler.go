package session

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/andersonjoseph/soundgo/internal/user"
)

type Entity struct {
	ID    string
	Token string
}

type Handler struct {
	repository     Repository
	userRepository user.Repository
	logger         *slog.Logger
	hasher         shared.PasswordHasher
}

func (h Handler) isPasswordValid(ctx context.Context, req *api.SessionInput) bool {
	u, err := h.userRepository.GetByUsername(ctx, req.Username)
	if err != nil {
		return false
	}

	return !h.hasher.Compare(u.Password, req.Password)
}

// POST /sessions
func (h Handler) CreateSession(ctx context.Context, req *api.SessionInput) (api.CreateSessionRes, error) {
	h.logger.Info(
		"creating session",
		slog.Group(
			"input",
			"username",
			req.Username,
		),
	)

	if !h.isPasswordValid(ctx, req) {
		return &api.Unauthorized{}, nil
	}

	session, err := h.repository.Save(req.Username, SaveInput{
		Token: "123",
	})

	if err != nil {
		h.logger.Error("error", slog.Group("error", "message", err.Error()))
		return nil, err
	}

	return &api.Session{
		Token: session.Token,
	}, nil
}

// DELETE /sessions
func (h Handler) DeleteSession(ctx context.Context) (api.DeleteSessionRes, error) {
	IDVal := ctx.Value("userID")

	ID, ok := IDVal.(string)
	if !ok {
		return nil, fmt.Errorf("bad ID")
	}

	h.logger.Info(
		"deleting session",
	)

	return &api.DeleteSessionNoContent{}, h.repository.Delete(ID)
}
