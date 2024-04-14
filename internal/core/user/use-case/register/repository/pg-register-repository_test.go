package repository_test

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/andersonjoseph/soundgo/internal/core/user/use-case/register/repository"
	"github.com/andersonjoseph/soundgo/internal/testhelper"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestIntegrationPgRepository_Save(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping integration tests")
	}

	connString := os.Getenv("DB")

	conn, err := pgxpool.New(context.Background(), connString)

	if err != nil {
		t.Fatal(err)
	}
	repo := repository.NewPGRepository(conn)

	type args struct {
		ctx context.Context
		dto repository.DtoSaveUser
	}
	tests := []struct {
		name    string
		args    args
		want    repository.DtoUser
		wantErr bool
	}{
		{
			name: "save user",
			args: args{
				ctx: context.TODO(),
				dto: testhelper.CreateDtoUser(t, "user@mail.com", "user", "password"),
			},
			want:    repository.DtoUser{Id: 1, Email: "user@mail.com", Username: "user", Password: "password"},
			wantErr: false,
		},
		{
			name: "save user with existing email",
			args: args{
				ctx: context.TODO(),
				dto: testhelper.CreateDtoUser(t, "user@mail.com", "otheruser", "password"),
			},
			want:    repository.DtoUser{},
			wantErr: true,
		},
		{
			name: "save user with existing username",
			args: args{
				ctx: context.TODO(),
				dto: testhelper.CreateDtoUser(t, "other@mail.com", "user", "password"),
			},
			want:    repository.DtoUser{},
			wantErr: true,
		},
		{
			name: "save another user",
			args: args{
				ctx: context.TODO(),
				dto: testhelper.CreateDtoUser(t, "other@mail.com", "otheruser", "password"),
			},
			want:    repository.DtoUser{Id: 2, Email: "other@mail.com", Username: "otheruser", Password: "password"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.Save(context.TODO(), tt.args.dto)

			tt.want.Id = got.Id

			if (err != nil) != tt.wantErr {
				t.Errorf("Save() test: %v, error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Save() test: %v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}

}
