package audio

import (
	"context"
	"errors"
	"testing"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/internaltest"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/andersonjoseph/soundgo/internal/user"
	"github.com/brianvoe/gofakeit/v7"
)

func TestSave(t *testing.T) {
	pool := internaltest.GetPgPool(t)
	userRepo := user.NewPGRepository(pool)

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
			name: "saving audio",
			args: args{
				ctx: context.TODO(),
				i: SaveInput{
					ID:          internaltest.GenerateUUID(t),
					Title:       gofakeit.BookTitle(),
					Description: gofakeit.Name(),
					UserID:      createRandomUser(t, userRepo).ID,
					Status:      api.AudioStatusPublished,
				},
			},
		},
		{
			name: "saving audio for a non existing user",
			args: args{
				ctx: context.TODO(),
				i: SaveInput{
					ID:          internaltest.GenerateUUID(t),
					Title:       gofakeit.BookTitle(),
					Description: gofakeit.Name(),
					UserID:      internaltest.GenerateUUID(t),
					Status:      api.AudioStatusPublished,
				},
			},
			err: shared.ErrNotFound,
		},
	}

	r := NewPGRepository(pool)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := r.Save(tt.args.ctx, tt.args.i)

			if !errors.Is(err, tt.err) {
				t.Errorf("Test failed: err expected: %v. received: %v", tt.err, err)
			}

			if tt.err != nil {
				return
			}

			if e.ID != tt.args.i.ID {
				t.Errorf("Test failed: ID expected: %v. received: %v", tt.args.i.ID, e.ID)
			}

			if e.Title != tt.args.i.Title {
				t.Errorf("Test failed: title expected: %v. received: %v", tt.args.i.Title, e.Title)
			}

			if e.Description != tt.args.i.Description {
				t.Errorf("Test failed: Description expected: %v. received: %v", tt.args.i.Description, e.Description)
			}

			if e.UserID != tt.args.i.UserID {
				t.Errorf("Test failed: UserID expected: %v. received: %v", tt.args.i.UserID, e.UserID)
			}

			if e.Status != tt.args.i.Status {
				t.Errorf("Test failed: Status expected: %v. received: %v", tt.args.i.Status, e.Status)
			}

			if e.CreatedAt.IsZero() {
				t.Errorf("Test failed: CreatedAt is Zero")
			}

			if e.Playcount != 0 {
				t.Errorf("Test failed: initial playcount is not Zero")
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
