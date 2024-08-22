package security

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type Handler struct {
	logger *slog.Logger
}

func NewHandler(logger *slog.Logger) Handler {
	return Handler{
		logger: logger,
	}
}

func (h Handler) GenerateToken(owner string) (string, error) {
	tok, err := jwt.NewBuilder().
		IssuedAt(time.Now()).
		Subject(owner).
		Build()

	if err != nil {
		h.logger.Error(
			"error building JWT",
			"msg",
			err.Error(),
			slog.Group("http_info", "status", 401),
		)
		return "", err
	}

	key, err := getJWTKey()
	if err != nil {
		h.logger.Error(
			"error getting JWT key",
			"msg",
			err.Error(),
			slog.Group("http_info", "status", 401),
		)
		return "", err
	}

	signedToken, err := jwt.Sign(tok, jwt.WithKey(jwa.HS512, []byte(key)))
	if err != nil {
		h.logger.Error(
			"error signing JWT",
			"msg",
			err.Error(),
			slog.Group("http_info", "status", 401),
		)

		return "", err
	}

	return string(signedToken), nil
}

func (h Handler) HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error) {
	key, err := getJWTKey()
	if err != nil {
		h.logger.Error(
			"error getting JWT key",
			"msg",
			err.Error(),
			slog.Group("http_info", "status", 401),
		)

		return nil, err
	}

	verifiedToken, err := jwt.Parse([]byte(t.Token), jwt.WithKey(jwa.HS512, []byte(key)))

	if err != nil {
		h.logger.Error(
			"error parsing JWT key",
			"msg",
			err.Error(),
			slog.Group("http_info", "status", 401),
		)

		return nil, err
	}

	ctx = context.WithValue(ctx, "session", verifiedToken.Subject())
	return ctx, nil
}

func getJWTKey() (string, error) {
	key, ok := os.LookupEnv("JWT_KEY")
	if !ok {
		return "", fmt.Errorf("JWT_KEY is not present in environment")
	}

	return key, nil
}
