package resetpassword

import (
	"context"
	"fmt"

	"github.com/andersonjoseph/soundgo/internal/core/user/model"
	"github.com/andersonjoseph/soundgo/internal/core/user/service"
	"github.com/andersonjoseph/soundgo/internal/core/user/use-case/reset-password/repository"
	"github.com/andersonjoseph/soundgo/internal/shared"
)

type useCase struct {
	repository     repository.ResetPasswordRepository
	tokenHandler   service.RequestPasswordTokenHandler
	passwordHasher shared.SecretHasher
}

func New(
	repo repository.ResetPasswordRepository,
	hasher shared.SecretHasher,
	tokenVerifier service.RequestPasswordTokenHandler,
) useCase {
	return useCase{
		repository:     repo,
		tokenHandler:   tokenVerifier,
		passwordHasher: hasher,
	}
}

type Dto struct {
	Token    string
	Password string
}

func (s useCase) ResetPassword(ctx context.Context, dto Dto) error {
	userID, err := s.tokenHandler.Decode(dto.Token)

	if err != nil {
		return err
	}

	passwordHash, err := s.repository.FindPasswordByID(ctx, userID)

	if err != nil {
		return err
	}

	password := model.NewPasswordFromHash(passwordHash, s.passwordHasher)

	if password.Compare(dto.Password) {
		return fmt.Errorf("new password can not be the same: %w", shared.ErrBadRequest)
	}

	newPassword, err := model.NewPassword(dto.Password, s.passwordHasher)

	if err != nil {
		return err
	}

	return s.repository.UpdatePassword(ctx, userID, newPassword)
}
