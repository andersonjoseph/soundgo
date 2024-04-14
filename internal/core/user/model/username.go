package model

import (
	"fmt"
	"regexp"

	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/go-playground/validator/v10"
)

type Username struct {
	value string
}

func (u Username) String() string {
	return u.value
}

var usernameValidator = (func() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		return regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(val)
	})

	return v
})()

func NewUsername(u string) (Username, error) {
	if err := usernameValidator.Var(u, "username,required,gt=3,lt=32"); err != nil {
		return Username{}, fmt.Errorf("username %w: %w", shared.ErrValidation, err)
	}

	return Username{
		value: u,
	}, nil
}
