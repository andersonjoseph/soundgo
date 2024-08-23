package shared

import "log/slog"

type FakeEmailSender struct {
	logger *slog.Logger
}

func NewFakeEmailSender(logger *slog.Logger) FakeEmailSender {
	return FakeEmailSender{
		logger: logger,
	}
}

func (s FakeEmailSender) SendPasswordCode(email string, code string) error {
	s.logger.Debug("sending password code", "email", email, "code", code)

	return nil
}
