package service

import (
	"log/slog"

	"github.com/andersonjoseph/soundgo/internal/core/user/model"
)

type fakePasswordTokenSender struct {
	logger slog.Logger
}

func NewFakePasswordTokenSender(logger slog.Logger) fakePasswordTokenSender {
	return fakePasswordTokenSender{
		logger: logger,
	}
}

func (ts fakePasswordTokenSender) Send(u model.User, tok string) error {
	ts.logger.Info("password token has been sent", "userID", u.ID, "token", tok)
	return nil
}
