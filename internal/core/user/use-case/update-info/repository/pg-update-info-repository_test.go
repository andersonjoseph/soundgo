package repository_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/andersonjoseph/soundgo/internal/core/user/model"
	"github.com/andersonjoseph/soundgo/internal/core/user/use-case/update-info/repository"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/andersonjoseph/soundgo/internal/testhelper"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestIntegrationPgRepository_UpdateInfo(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests")
	}

	connString := os.Getenv("DB")

	conn, err := pgxpool.New(context.Background(), connString)

	if err != nil {
		t.Fatal(err)
	}

	repo := repository.NewPGRepository(conn)

	createdUser, err := testhelper.CreateUser(t, conn, "user_to_update@mail.com", "user_to_update", "password")
	otherUser, err := testhelper.CreateUser(t, conn, "existing_user_update@mail.com", "existing_user_update", "password")

	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		ctx    context.Context
		params repository.UpdateInfoParams
	}

	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "update username",
			args: args{
				ctx: context.TODO(),
				params: repository.UpdateInfoParams{
					ID: createdUser.ID,
					Username: func() model.Username {
						u, err := model.NewUsername("new_username")

						if err != nil {
							t.Fatal("err")
						}

						return u
					}(),
				},
			},
			want: nil,
		},
		{
			name: "update nonexisting user",
			args: args{
				ctx: context.TODO(),
				params: repository.UpdateInfoParams{
					ID: 69429,
					Username: func() model.Username {
						u, err := model.NewUsername("new_username1")

						if err != nil {
							t.Fatal("err")
						}

						return u
					}(),
				},
			},
			want: shared.ErrNotFound,
		},
		{
			name: "update taken username",
			args: args{
				ctx: context.TODO(),
				params: repository.UpdateInfoParams{
					ID:       createdUser.ID,
					Username: otherUser.Username,
				},
			},
			want: shared.ErrRecordAlreadyExists,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.UpdateInfo(context.TODO(), tt.args.params)

			t.Log(err)

			if !errors.Is(err, tt.want) {
				t.Errorf("UpdateInfo() test: %v, error = %v, want %v", tt.name, err, tt.want)
				return
			}
		})
	}
}
