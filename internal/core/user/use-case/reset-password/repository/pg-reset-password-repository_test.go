package repository_test

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/andersonjoseph/soundgo/internal/core/user/model"
	"github.com/andersonjoseph/soundgo/internal/core/user/use-case/reset-password/repository"
	"github.com/andersonjoseph/soundgo/internal/testhelper"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestIntegrationPgRepository_FindPasswordByID(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests")
	}

	connString := os.Getenv("DB")

	conn, err := pgxpool.New(context.Background(), connString)
	repo := repository.NewPGRepository(conn)

	if err != nil {
		t.Fatal(err)
	}

	createdUser, err := testhelper.CreateUser(t, conn, "user_resetting_password@mail.com", "user_resetting_password", "password")

	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		ctx context.Context
		id  int
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "find user password",
			args: args{
				ctx: context.TODO(),
				id:  createdUser.ID,
			},
			want:    createdUser.Password.String(),
			wantErr: false,
		},
		{
			name: "find nonexisting password",
			args: args{
				ctx: context.TODO(),
				id:  69420,
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.FindPasswordByID(context.TODO(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindPasswordByID() test: %v, error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindPasswordByID() test: %v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestIntegrationPgRepository_UpdatePassword(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests")
	}

	connString := os.Getenv("DB")

	conn, err := pgxpool.New(context.Background(), connString)
	repo := repository.NewPGRepository(conn)

	if err != nil {
		t.Fatal(err)
	}

	createdUser, err := testhelper.CreateUser(t, conn, "user_updating_password@mail.com", "user_updating_password", "password")

	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		ctx      context.Context
		id       int
		password model.Password
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "update user password",
			args: args{
				ctx:      context.TODO(),
				id:       createdUser.ID,
				password: testhelper.CreateDtoUser(t, createdUser.Email.String(), createdUser.Username.String(), "1234567890").Password,
			},
			wantErr: false,
		},
		{
			name: "update nonexisting user password",
			args: args{
				ctx:      context.TODO(),
				id:       69420,
				password: testhelper.CreateDtoUser(t, createdUser.Email.String(), createdUser.Username.String(), "1234567890").Password,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.UpdatePassword(context.TODO(), tt.args.id, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindPasswordByID() test: %v, error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
		})
	}
}
