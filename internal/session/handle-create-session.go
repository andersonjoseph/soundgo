package session // POST /sessions

import (
	"context"
	"errors"
	"fmt"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/andersonjoseph/soundgo/internal/user"
)

func (h Handler) CreateSession(ctx context.Context, req *api.SessionInput) (api.CreateSessionRes, error) {
	u, err := h.userRepository.GetByUsername(ctx, req.Username)
	if errors.Is(err, shared.ErrNotFound) {
		return &api.Unauthorized{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error while getting session by username: %w", err)
	}

	if !h.hasher.Compare(u.Password, req.Password) {
		return &api.Unauthorized{}, nil
	}

	input, err := h.createSaveInput(u)
	if err != nil {
		return nil, fmt.Errorf("error while creating save input: %w", err)
	}

	session, err := h.repository.Save(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("error while saving session: %w", err)
	}

	return &api.Session{
		Token: session.Token,
	}, nil
}

func (h Handler) createSaveInput(u user.Entity) (SaveInput, error) {
	ID, err := shared.GenerateUUID()
	if err != nil {
		return SaveInput{}, err
	}

	token, err := h.jwtHandler.GenerateToken(ID)
	if err != nil {
		return SaveInput{}, err
	}

	return SaveInput{
		ID:     ID,
		UserID: u.ID,
		Token:  token,
	}, nil
}
