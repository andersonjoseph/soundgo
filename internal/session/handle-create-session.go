package session // POST /sessions

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/andersonjoseph/soundgo/internal/user"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func (h Handler) CreateSession(ctx context.Context, req *api.SessionInput) (api.CreateSessionRes, error) {
	h.logger.Info(
		"creating session",
		slog.Group(
			"input",
			"username",
			req.Username,
		),
	)
	u, err := h.userRepository.GetByUsername(ctx, req.Username)
	if err != nil {
		switch {
		case errors.Is(err, shared.ErrNotFound):
			h.handleError(shared.ErrUnauthorized)
		default:
			return h.handleError(fmt.Errorf("error while getting user by username: %w", err))
		}
	}

	if !h.hasher.Compare(u.Password, req.Password) {
		return &api.Unauthorized{}, nil
	}

	input, err := createSaveInput(u)
	if err != nil {
		return h.handleError(fmt.Errorf("error while creating save input: %w", err))
	}

	session, err := h.repository.Save(ctx, input)
	if err != nil {
		return h.handleError(fmt.Errorf("error while saving session: %w", err))
	}

	h.logger.Info("session created", "ID", session.ID)
	return &api.Session{
		Token: session.Token,
	}, nil
}

func createSaveInput(u user.Entity) (SaveInput, error) {
	ID, err := shared.GenerateUUID()
	if err != nil {
		return SaveInput{}, err
	}

	token, err := createToken(ID)
	if err != nil {
		return SaveInput{}, err
	}

	return SaveInput{
		ID:     ID,
		UserID: u.ID,
		Token:  token,
	}, nil
}

func createToken(owner string) (string, error) {
	tok, err := jwt.NewBuilder().
		IssuedAt(time.Now()).
		Subject(owner).
		Build()

	if err != nil {
		return "", fmt.Errorf("error building JWT: %w", err)
	}

	key, ok := os.LookupEnv("JWT_KEY")
	if !ok {
		panic("JWT_KEY is not present in environment")
	}

	signedToken, err := jwt.Sign(tok, jwt.WithKey(jwa.HS512, []byte(key)))
	if err != nil {
		return "", fmt.Errorf("error signing JWT: %w", err)
	}

	return string(signedToken), nil
}
