package password

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/andersonjoseph/soundgo/internal/internaltest"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/andersonjoseph/soundgo/internal/user"
	"github.com/brianvoe/gofakeit/v7"
)

func TestSaveResetRequest(t *testing.T) {
	pool := internaltest.GetPgPool(t)
	userRepo := user.NewPGRepository(pool)
	repo := NewPGRepository(pool)

	type args struct {
		ctx context.Context
		i   SaveResetRequestInput
	}

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "saving reset request",
			args: args{
				ctx: context.TODO(),
				i: SaveResetRequestInput{
					userID:    createRandomUser(t, userRepo).ID,
					code:      "123",
					expiresAt: time.Now(),
				},
			},
			err: nil,
		},
		{
			name: "saving reset request for a non existing user returns error",
			args: args{
				ctx: context.TODO(),
				i: SaveResetRequestInput{
					userID:    internaltest.GenerateUUID(t),
					code:      "123",
					expiresAt: time.Now(),
				},
			},
			err: shared.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.SaveResetRequest(tt.args.ctx, tt.args.i)

			if !errors.Is(err, tt.err) {
				t.Errorf("Test failed: wanted:%v received:%v ", tt.err, err)
			}
		})

	}
}

func TestGetRequestByUserID(t *testing.T) {
	pool := internaltest.GetPgPool(t)
	userRepo := user.NewPGRepository(pool)
	repo := NewPGRepository(pool)

	user := createRandomUser(t, userRepo)

	err := repo.SaveResetRequest(context.TODO(), SaveResetRequestInput{
		userID:    user.ID,
		code:      "123",
		expiresAt: time.Now(),
	})

	if err != nil {
		t.Fatalf("Test failed: error occured while creating reset request. received: %v", err)
	}

	type args struct {
		ctx    context.Context
		userID string
	}

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "saving reset request",
			args: args{
				ctx:    context.TODO(),
				userID: user.ID,
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := repo.GetRequestByUserID(tt.args.ctx, tt.args.userID)

			if !errors.Is(err, tt.err) {
				t.Errorf("Test failed: wanted:%v received:%v ", tt.err, err)
			}

			if req.Code != "123" {
				t.Errorf("Test failed: recevied code is not equal to 123. received:%v ", req.Code)
			}

			if req.UserID != tt.args.userID {
				t.Errorf("Test failed: recevied userID is not equal to expected. received:%v. wanted:%v ", req.Code, tt.args.userID)
			}

			if req.ExpiresAt.IsZero() {
				t.Errorf("Test failed: expiresAt is zero")
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
