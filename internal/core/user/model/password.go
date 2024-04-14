package model

import (
	"fmt"

	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/go-playground/validator/v10"
)

type Password struct {
	value  string
	hasher shared.SecretHasher
}

func (p Password) String() string {
	return p.value
}

var passwordValidator = validator.New()

func NewPassword(p string, hasher shared.SecretHasher) (Password, error) {
	if err := usernameValidator.Var(p, "required,min=8,max=256"); err != nil {
		return Password{}, fmt.Errorf("password %w: %w", shared.ErrValidation, err)
	}

	hashedPwd, err := hasher.Hash(p)
	if err != nil {
		return Password{}, fmt.Errorf("password %w:", err)
	}

	return Password{
		value:  hashedPwd,
		hasher: hasher,
	}, nil
}

func NewPasswordFromHash(p string, hasher shared.SecretHasher) Password {
	return Password{
		value:  p,
		hasher: hasher,
	}
}

func (p Password) Compare(plainPassword string) bool {
	return p.hasher.Compare(p.value, plainPassword)
}
