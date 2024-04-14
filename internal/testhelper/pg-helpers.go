package testhelper

import (
	"context"
	"testing"

	"github.com/andersonjoseph/soundgo/internal/core/user/model"
	"github.com/andersonjoseph/soundgo/internal/core/user/use-case/register"
	"github.com/andersonjoseph/soundgo/internal/core/user/use-case/register/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type testHasher struct{}

func (h testHasher) Hash(p string) (string, error) {
	return p, nil
}

func (h testHasher) Compare(hp string, sp string) bool {
	return hp == sp
}

var hasher = testHasher{}

func CreateDtoUser(t *testing.T, email, username, password string) repository.DtoSaveUser {
	t.Helper()

	e, err := model.NewEmail(email)

	if err != nil {
		t.Fatal(err)
	}

	u, err := model.NewUsername(username)

	if err != nil {
		t.Fatal(err)
	}

	p, err := model.NewPassword(password, hasher)

	if err != nil {
		t.Fatal(err)
	}

	return repository.DtoSaveUser{
		Email:    e,
		Username: u,
		Password: p,
	}
}

func CreateUser(t *testing.T, conn *pgxpool.Pool, email, username, password string) (model.User, error) {
	t.Helper()

	repo := repository.NewPGRepository(conn)

	service := register.New(repo, hasher)

	return service.RegisterUser(context.TODO(), register.Dto{
		Email:    email,
		Username: username,
		Password: password,
	})
}
