package repository_test

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/andersonjoseph/soundgo/internal/core/user/model"
	"github.com/andersonjoseph/soundgo/internal/core/user/use-case/login/repository"
	"github.com/andersonjoseph/soundgo/internal/testhelper"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestIntegrationPgRepository_FindPasswordByUsername(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests")
	}

	connString := os.Getenv("DB")

	conn, err := pgxpool.New(context.Background(), connString)

	if err != nil {
		t.Fatal(err)
	}

	repo := repository.NewPGRepository(conn)

	createdUser, err := testhelper.CreateUser(t, conn, "find_password_user@mail.com", "find_password_user", "password")

	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		ctx      context.Context
		username model.Username
	}

	tests := []struct {
		name    string
		args    args
		want    repository.DtoPassword
		wantErr bool
	}{
		{
			name: "find user password",
			args: args{
				ctx:      context.TODO(),
				username: createdUser.Username,
			},
			want: repository.DtoPassword{
				UserID: createdUser.ID,
				Value:  createdUser.Password.String(),
			},
			wantErr: false,
		},
		{
			name: "find nonexisting user password",
			args: args{
				ctx:      context.TODO(),
				username: testhelper.CreateDtoUser(t, "nonexisting_user@mail.com", "nonexistinuser", "password").Username,
			},
			want:    repository.DtoPassword{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.FindPasswordByUsername(context.TODO(), tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByPassword() test: %v, error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByPassword() test: %v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}

}

func TestIntegrationPgRepository_SaveSession(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests")
	}

	connString := os.Getenv("DB")

	conn, err := pgxpool.New(context.Background(), connString)
	repo := repository.NewPGRepository(conn)

	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		ctx context.Context
		dto repository.DtoCreateSession
	}

	createdUser, err := testhelper.CreateUser(t, conn, "save_session_user@mail.com", "save_session_user", "password")

	tests := []struct {
		name    string
		args    args
		want    repository.DtoSession
		wantErr bool
	}{
		{
			name: "create session",
			args: args{
				ctx: context.TODO(),
				dto: repository.DtoCreateSession{
					UserID:       createdUser.ID,
					CreationDate: time.Now().Format(time.RFC3339),
				},
			},
			wantErr: false,
		},
		{
			name: "create session for a nonexisting user",
			args: args{
				ctx: context.TODO(),
				dto: repository.DtoCreateSession{
					UserID:       42069,
					CreationDate: time.Now().Format(time.RFC3339),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.SaveSession(tt.args.ctx, tt.args.dto)

			tt.want.ID = got.ID
			tt.want.CreationDate = got.CreationDate

			if (err != nil) != tt.wantErr {
				t.Errorf("SaveSession() test: %v, error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SaveSession() test: %v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
