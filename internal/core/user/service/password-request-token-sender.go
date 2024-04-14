package service

import "github.com/andersonjoseph/soundgo/internal/core/user/model"

type PasswordTokenSender interface {
	Send(u model.User, tok string) error
}
