package requestresetpassword

import (
	"context"

	"github.com/andersonjoseph/soundgo/internal/core/user/model"
	"github.com/andersonjoseph/soundgo/internal/core/user/service"
	"github.com/andersonjoseph/soundgo/internal/core/user/use-case/request-password-reset/repository"
)

type useCase struct {
	repo           repository.RequestResetPasswordRepository
	tokenGenerator service.RequestPasswordTokenHandler
	tokenSender    service.PasswordTokenSender
}

type Dto struct {
	Username string
	Email    string
}

func New(r repository.RequestResetPasswordRepository, tg service.RequestPasswordTokenHandler, ts service.PasswordTokenSender) useCase {
	return useCase{
		repo:           r,
		tokenGenerator: tg,
		tokenSender:    ts,
	}
}

func (s useCase) SendRequestPasswordToken(ctx context.Context, dto Dto) error {
	u, err := s.findUser(ctx, dto.Username, dto.Email)

	if err != nil {
		return err
	}

	tok, err := s.tokenGenerator.Generate(u.ID)

	if err != nil {
		return err
	}

	return s.tokenSender.Send(u, tok)
}

func (s useCase) findUser(ctx context.Context, username, email string) (model.User, error) {
	u, err := model.NewUsername(username)

	if err != nil {
		return model.User{}, err
	}

	e, err := model.NewEmail(email)

	if err != nil {
		return model.User{}, err
	}

	userDTO, err := s.repo.FindByUsernameAndEmail(ctx, u, e)

	if err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:       userDTO.ID,
		Username: u,
		Email:    e,
	}, nil
}
