package repository_test

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/andersonjoseph/soundgo/internal/core/user/model"
	"github.com/andersonjoseph/soundgo/internal/core/user/use-case/request-password-reset/repository"
	"github.com/andersonjoseph/soundgo/internal/testhelper"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestIntegrationPgRepository_FindUsernameAndEmail(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests")
	}

	connString := os.Getenv("DB")

	conn, err := pgxpool.New(context.Background(), connString)
	repo := repository.NewPGRepository(conn)

	if err != nil {
		t.Fatal(err)
	}

	createdUser, err := testhelper.CreateUser(t, conn, "req_res_password_user_exists@mail.com", "req_res_password_user_exists", "password")

	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		ctx      context.Context
		username model.Username
		email    model.Email
	}

	tests := []struct {
		name    string
		args    args
		want    repository.FindByUserDTO
		wantErr bool
	}{
		{
			name: "find user",
			args: args{
				ctx:      context.TODO(),
				username: createdUser.Username,
				email:    createdUser.Email,
			},
			want: repository.FindByUserDTO{
				ID:       createdUser.ID,
				Username: createdUser.Username.String(),
				Email:    createdUser.Email.String(),
			},
			wantErr: false,
		},
		{
			name: "find nonexisting user",
			args: args{
				ctx: context.TODO(),
				username: (func() model.Username {
					u, err := model.NewUsername("req_res_pass_user_false")

					if err != nil {
						t.Fatal(err)
					}

					return u
				})(),
				email: (func() model.Email {
					u, err := model.NewEmail("req_res_pass_user_false@mail.com")

					if err != nil {
						t.Fatal(err)
					}

					return u
				})(),
			},
			want:    repository.FindByUserDTO{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.FindByUsernameAndEmail(context.TODO(), tt.args.username, tt.args.email)

			if (err != nil) != tt.wantErr {
				t.Errorf("FindByUsernameAndEmail(context.TODO(() test: %v, error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByUsernameAndEmail() test: %v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
