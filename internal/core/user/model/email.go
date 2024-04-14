package model

import (
	"fmt"
	"strings"

	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/go-playground/validator/v10"
)

type Email struct {
	value string
}

var emailValidator = validator.New()

func NewEmail(e string) (Email, error) {
	e = strings.ToLower(e)
	if err := emailValidator.Var(e, "required,email"); err != nil {
		return Email{}, fmt.Errorf("email %w: %w", shared.ErrValidation, err)
	}

	return Email{
		value: e,
	}, nil
}

func (e Email) String() string {
	return e.value
}
