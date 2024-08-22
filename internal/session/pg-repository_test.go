package session

import (
	"context"
	"errors"
	"testing"

	"github.com/andersonjoseph/soundgo/internal/internaltest"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/andersonjoseph/soundgo/internal/user"
	"github.com/brianvoe/gofakeit/v7"
)

func TestSave(t *testing.T) {
	pool := internaltest.GetPgPool(t)
	userRepo := user.NewPGRepository(pool)
	user := createRandomUser(t, userRepo)

	type args struct {
		ctx context.Context
		i   SaveInput
	}

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "saving session",
			args: args{
				ctx: context.TODO(),
				i: SaveInput{
					ID:     internaltest.GenerateUUID(t),
					Token:  "123",
					UserID: user.ID,
				},
			},
		},
		{
			name: "saving multiple session for the same user",
			args: args{
				ctx: context.TODO(),
				i: SaveInput{
					ID:     internaltest.GenerateUUID(t),
					Token:  "123",
					UserID: user.ID,
				},
			},
		},
		{
			name: "saving session with for a non existing userID returns an error",
			args: args{
				ctx: context.TODO(),
				i: SaveInput{
					ID:     internaltest.GenerateUUID(t),
					Token:  "123",
					UserID: internaltest.GenerateUUID(t),
				},
			},
			err: shared.ErrNotFound,
		},
	}

	r := NewPGRepository(pool)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := r.Save(tt.args.ctx, tt.args.i)

			if errors.Is(err, tt.err) {
				return
			}

			if tt.err != err {
				t.Errorf("Test failed: err expected %v. received: %v", tt.err, err)
			}

			if e.ID != tt.args.i.ID {
				t.Errorf("received ID is not equal to expected. received='%s' expected='%s'", e.ID, tt.args.i.ID)
			}

			if e.Token != tt.args.i.Token {
				t.Errorf("received Token is not equal to expected. received='%s' expected='%s'", e.Token, tt.args.i.Token)
			}
		})
	}
}

func createRandomUser(t *testing.T, r user.PGRepository) user.Entity {
	t.Helper()

	u, err := r.Save(context.TODO(), user.SaveInput{
		ID:       internaltest.GenerateUUID(t),
		Username: gofakeit.Username(),
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, true, true, 8),
	})

	if err != nil {
		t.Fatalf("Test failed: error occured while creating test user. received: %v", err)
	}

	return u
}
