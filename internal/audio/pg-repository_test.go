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
					Status:      api.AudioInputMultipartStatusPublished,
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
					Status:      api.AudioInputMultipartStatusPublished,
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

			expectedStatus, err := e.Status.MarshalText()
			if err != nil {
				t.Fatalf("Test failed: error while marshalling expected status: %v", err)
			}
			receviedStatus, err := tt.args.i.Status.MarshalText()
			if err != nil {
				t.Fatalf("Test failed: error while marshalling expected status: %v", err)
			}

			if string(expectedStatus) != string(receviedStatus) {
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

func TestGet(t *testing.T) {
	pool := internaltest.GetPgPool(t)
	userRepo := user.NewPGRepository(pool)
	repo := NewPGRepository(pool)

	audio := createRandomAudio(t, repo, userRepo)

	type args struct {
		ctx context.Context
		id  string
	}

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "getting audio",
			args: args{
				ctx: context.TODO(),
				id:  audio.ID,
			},
		},
		{
			name: "getting non existing audio",
			args: args{
				ctx: context.TODO(),
				id:  internaltest.GenerateUUID(t),
			},
			err: shared.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := repo.Get(tt.args.ctx, tt.args.id)

			if !errors.Is(err, tt.err) {
				t.Errorf("Test failed: err expected: %v. received: %v", tt.err, err)
			}

			if tt.err != nil {
				return
			}

			if e.ID != audio.ID {
				t.Errorf("Test failed: ID expected: %v. received: %v", audio.ID, e.ID)
			}

			if e.Title != audio.Title {
				t.Errorf("Test failed: title expected: %v. received: %v", audio.Title, e.Title)
			}

			if e.Description != audio.Description {
				t.Errorf("Test failed: Description expected: %v. received: %v", audio.Description, e.Description)
			}

			if e.UserID != audio.UserID {
				t.Errorf("Test failed: UserID expected: %v. received: %v", audio.UserID, e.UserID)
			}

			if e.Status != audio.Status {
				t.Errorf("Test failed: Status expected: %v. received: %v", audio.Status, e.Status)
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

func createRandomAudio(t *testing.T, repo PGRepository, userRepo user.PGRepository) Entity {
	t.Helper()

	a, err := repo.Save(context.TODO(), SaveInput{
		ID:          internaltest.GenerateUUID(t),
		Title:       gofakeit.BookTitle(),
		Description: gofakeit.Name(),
		UserID:      createRandomUser(t, userRepo).ID,
		Status:      api.AudioInputMultipartStatusPublished,
	})

	if err != nil {
		t.Fatalf("Test failed: error occured while creating test user. received: %v", err)
	}

	return a
}
