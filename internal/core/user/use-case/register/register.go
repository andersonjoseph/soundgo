package register

import (
	"context"

	"github.com/andersonjoseph/soundgo/internal/core/user/model"
	"github.com/andersonjoseph/soundgo/internal/core/user/use-case/register/repository"
	"github.com/andersonjoseph/soundgo/internal/shared"
)

type useCase struct {
	repository     repository.UserRepository
	passwordHasher shared.SecretHasher
}

func New(repo repository.UserRepository, hasher shared.SecretHasher) useCase {
	return useCase{
		repository:     repo,
		passwordHasher: hasher,
	}
}

type Dto struct {
	Email    string
	Username string
	Password string
}

func (s useCase) RegisterUser(ctx context.Context, dto Dto) (model.User, error) {
	userToSave, err := s.mapToDtoSaveUser(dto)

	if err != nil {
		return model.User{}, err
	}

	savedUser, err := s.repository.Save(ctx, userToSave)

	if err != nil {
		return model.User{}, err
	}

	u, err := s.mapFromUserRepository(savedUser)

	if err != nil {
		return model.User{}, err
	}

	return u, nil
}

func (s useCase) mapFromUserRepository(dto repository.DtoUser) (model.User, error) {
	e, err := model.NewEmail(dto.Email)
	if err != nil {
		return model.User{}, err
	}

	u, err := model.NewUsername(dto.Username)
	if err != nil {
		return model.User{}, err
	}

	p, err := model.NewPassword(dto.Password, s.passwordHasher)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:       dto.Id,
		Username: u,
		Email:    e,
		Password: p,
	}, nil
}

func (s useCase) mapToDtoSaveUser(dto Dto) (repository.DtoSaveUser, error) {
	email, err := model.NewEmail(dto.Email)
	if err != nil {
		return repository.DtoSaveUser{}, err
	}

	username, err := model.NewUsername(dto.Username)
	if err != nil {
		return repository.DtoSaveUser{}, err
	}

	password, err := model.NewPassword(dto.Password, s.passwordHasher)
	if err != nil {
		return repository.DtoSaveUser{}, err
	}

	return repository.DtoSaveUser{
		Email:    email,
		Username: username,
		Password: password,
	}, nil
}
