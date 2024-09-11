package password

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/andersonjoseph/soundgo/internal/user"
)

type Handler struct {
	userRepository user.Repository
	repository     Repository
	hasher         shared.PasswordHasher
	emailSender    shared.EmailSender
}

type RequestReset struct {
	ID        uint64
	UserID    string
	Code      string
	ExpiresAt time.Time
}

func NewHandler(
	repo Repository,
	userRepo user.Repository,
	hasher shared.PasswordHasher,
	emailSender shared.EmailSender,
) Handler {
	return Handler{
		repository:     repo,
		userRepository: userRepo,
		hasher:         hasher,
		emailSender:    emailSender,
	}
}

// POST /password-reset
func (h Handler) CreatePasswordResetRequest(ctx context.Context, req *api.PasswordResetRequestInput) (api.CreatePasswordResetRequestRes, error) {
	user, err := h.userRepository.GetByEmail(ctx, req.Email)

	if errors.Is(err, shared.ErrNotFound) {
		return &api.CreatePasswordResetRequestNoContent{}, nil
	}

	if err != nil {
		return nil, err
	}

	code, err := generateCode()
	if err != nil {
		return nil, err
	}

	if err = h.repository.SaveResetRequest(ctx, SaveResetRequestInput{
		userID:    user.ID,
		code:      code,
		expiresAt: time.Now().Add(time.Minute * 30),
	}); err != nil {
		return nil, err
	}

	if err := h.emailSender.SendPasswordCode(req.Email, code); err != nil {
		return nil, err
	}

	return &api.CreatePasswordResetRequestNoContent{}, nil
}

func generateCode() (string, error) {
	b := make([]byte, 3)
	_, err := rand.Read(b)

	if err != nil {
		return "", fmt.Errorf("error while generating password code: %w", err)
	}

	return fmt.Sprintf("%x", b), err
}

// PUT /password-reset
func (h Handler) ResetPassword(ctx context.Context, req *api.PasswordResetInput) (api.ResetPasswordRes, error) {
	u, err := h.userRepository.GetByEmail(ctx, req.Email)

	if errors.Is(err, shared.ErrNotFound) {
		return &api.ResetPasswordBadRequest{}, err
	}
	if err != nil {
		return nil, err
	}

	resetRequest, err := h.repository.GetRequestByUserID(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	if resetRequest.Code != req.Code {
		return &api.ResetPasswordBadRequest{}, err
	}

	hashedPassword, err := h.hasher.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	_, err = h.userRepository.Update(ctx, u.ID, user.UpdateInput{Password: hashedPassword})

	return &api.ResetPasswordNoContent{}, err
}
