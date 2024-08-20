package password

import (
	"context"
	"log/slog"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/andersonjoseph/soundgo/internal/user"
)

type Handler struct {
	userRepository user.Repository
	repository     Repository
	logger         *slog.Logger
	hasher         shared.PasswordHasher
	emailSender    shared.EmailSender
}

// POST /password-reset
func (h Handler) CreatePasswordResetRequest(ctx context.Context, req *api.PasswordResetRequestInput) (api.CreatePasswordResetRequestRes, error) {
	h.logger.Info("requesting password reset",
		slog.Group(
			"input",
			"email",
			req.Email,
		),
	)
	userExists, err := h.userRepository.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !userExists {
		return &api.CreatePasswordResetRequestNoContent{}, nil
	}

	if err := h.emailSender.SendPasswordCode(req.Email); err != nil {
		return nil, err
	}

	return &api.CreatePasswordResetRequestNoContent{}, nil
}

// PUT /password-reset
func (h Handler) ResetPassword(ctx context.Context, req *api.PasswordResetInput) (api.ResetPasswordRes, error) {
	code, err := h.repository.GetCode(req.Email)
	if err != nil {
		return nil, err
	}

	if code != req.Code {
		return &api.ResetPasswordBadRequest{}, err
	}

	hashedPassword, err := h.hasher.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	return &api.ResetPasswordNoContent{}, h.repository.Save(req.Email, hashedPassword)
}
