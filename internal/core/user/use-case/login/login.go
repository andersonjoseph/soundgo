package login

import (
	"context"
	"errors"
	"time"

	"github.com/andersonjoseph/soundgo/internal/core/user/model"
	"github.com/andersonjoseph/soundgo/internal/core/user/use-case/login/repository"
	"github.com/andersonjoseph/soundgo/internal/shared"
)

type Dto struct {
	Username string
	Password string
}

type useCase struct {
	repository     repository.LoginRepository
	passwordHasher shared.SecretHasher
}

func New(repo repository.LoginRepository, hasher shared.SecretHasher) useCase {
	return useCase{
		repository:     repo,
		passwordHasher: hasher,
	}
}

func (s useCase) LoginUser(ctx context.Context, dto Dto) (model.Session, error) {
	userPasswordHash, err := s.findPasswordByUsername(ctx, dto.Username)

	if errors.Is(err, shared.ErrNotFound) {
		return model.Session{}, shared.ErrUnauthorized
	}

	if err != nil {
		return model.Session{}, err
	}

	userPassword := model.NewPasswordFromHash(userPasswordHash.Value, s.passwordHasher)

	if !userPassword.Compare(dto.Password) {
		return model.Session{}, shared.ErrUnauthorized
	}

	return s.createSession(ctx, repository.DtoCreateSession{
		UserID:       userPasswordHash.UserID,
		CreationDate: time.Now().Format(time.RFC3339),
	})
}

func (s useCase) findPasswordByUsername(ctx context.Context, username string) (repository.DtoPassword, error) {
	u, err := model.NewUsername(username)

	if err != nil {
		return repository.DtoPassword{}, err
	}

	return s.repository.FindPasswordByUsername(ctx, u)
}

func (s useCase) createSession(ctx context.Context, dto repository.DtoCreateSession) (model.Session, error) {
	sessionDTO, err := s.repository.SaveSession(ctx, dto)

	if err != nil {
		return model.Session{}, err
	}

	return model.Session{ID: sessionDTO.ID}, nil
}
