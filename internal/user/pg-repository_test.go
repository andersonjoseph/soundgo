package user

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/andersonjoseph/soundgo/internal/internaltest"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/brianvoe/gofakeit/v7"
)

func createRandomUser(t *testing.T, r PGRepository) Entity {
	t.Helper()

	u, err := r.Save(context.TODO(), SaveInput{
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

func TestSave(t *testing.T) {
	pool := internaltest.GetPgPool(t)

	type args struct {
		ctx context.Context
		i   SaveInput
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "save user",
			args: args{
				ctx: context.TODO(),
				i: SaveInput{
					ID:       internaltest.GenerateUUID(t),
					Username: gofakeit.Username(),
					Email:    gofakeit.Email(),
					Password: gofakeit.Password(true, true, true, true, true, 8),
				},
			},
			wantErr: false,
		},
		{
			name: "saving user with invalid uuid returns error",
			args: args{
				ctx: context.TODO(),
				i: SaveInput{
					ID:       "123",
					Username: gofakeit.Username(),
					Email:    gofakeit.Email(),
					Password: gofakeit.Password(true, true, true, true, true, 8),
				},
			},
			wantErr: true,
		},
		{
			name: "saving user with an username with length greater than 32 returns error",
			args: args{
				ctx: context.TODO(),
				i: SaveInput{
					ID:       internaltest.GenerateUUID(t),
					Username: "superduperlongusername3322",
					Email:    gofakeit.Email(),
					Password: gofakeit.Password(true, true, true, true, true, 8),
				},
			},
			wantErr: true,
		},
	}

	r := NewPGRepository(pool)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := r.Save(tt.args.ctx, tt.args.i)

			if tt.wantErr && err != nil {
				return
			}

			if tt.wantErr != (err != nil) {
				t.Errorf("Test failed: err was not expected. received: %v", err)
			}

			if e.CreatedAt.IsZero() {
				t.Errorf("CreatedAt is zero")
			}

			if e.Username != tt.args.i.Username {
				t.Errorf("received Username is not equal to expected. received='%s' expected='%s'", e.Username, tt.args.i.Username)
			}

			if e.Email != tt.args.i.Email {
				t.Errorf("received Username is not equal to expected. received='%s' expected='%s'", e.Password, tt.args.i.Password)
			}

			if e.Password != tt.args.i.Password {
				t.Errorf("received Password is not equal to expected. received='%s' expected='%s'", e.Password, tt.args.i.Password)
			}
		})
	}

	t.Run("duplicated record returns error", func(t *testing.T) {
		_, err := r.Save(tests[0].args.ctx, tests[0].args.i)

		if !errors.Is(err, shared.ErrAlreadyExists) {
			t.Errorf("Test failed: ErrAlreadyExists was expected. received: %v", err)
		}
	})
}

func TestUpdate(t *testing.T) {
	r := NewPGRepository(internaltest.GetPgPool(t))

	type args struct {
		ctx context.Context
		ID  string
		i   UpdateInput
	}

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "updating user",
			args: args{
				ctx: context.TODO(),
				ID:  createRandomUser(t, r).ID,
				i: UpdateInput{
					Username: gofakeit.Username(),
				},
			},
		},
		{
			name: "updating user with an existing username, returns an error",
			args: args{
				ctx: context.TODO(),
				ID:  createRandomUser(t, r).ID,
				i: UpdateInput{
					Username: createRandomUser(t, r).Username,
				},
			},
			err: shared.ErrAlreadyExists,
		},
		{
			name: "updating user with a non existing ID, returns an error",
			args: args{
				ctx: context.TODO(),
				ID:  internaltest.GenerateUUID(t),
				i: UpdateInput{
					Username: gofakeit.Username(),
				},
			},
			err: shared.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := r.Update(tt.args.ctx, tt.args.ID, tt.args.i)

			if errors.Is(err, tt.err) {
				return
			}

			if tt.err != err {
				t.Errorf("Test failed: err expected %v. received: %v", tt.err, err)
			}

			if e.Username != tt.args.i.Username {
				t.Errorf("received username is not equal to expected. received='%s' expected='%s'", e.Username, tt.args.i.Username)
			}
		})
	}
}

func TestGetByUsername(t *testing.T) {
	r := NewPGRepository(internaltest.GetPgPool(t))

	type args struct {
		ctx      context.Context
		username string
	}

	expectedUser := createRandomUser(t, r)

	tests := []struct {
		name string
		args args
		want Entity
		err  error
	}{
		{
			name: "getting user",
			args: args{
				ctx:      context.TODO(),
				username: expectedUser.Username,
			},
			want: expectedUser,
		},
		{
			name: "getting a user does not exist returns not found",
			args: args{
				ctx:      context.TODO(),
				username: "johndoe",
			},
			want: expectedUser,
			err:  shared.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := r.GetByUsername(tt.args.ctx, tt.args.username)

			if errors.Is(err, tt.err) {
				return
			}

			if tt.err != err {
				t.Errorf("Test failed: err expected %v. received: %v", tt.err, err)
			}

			if !reflect.DeepEqual(expectedUser, user) {
				t.Errorf("found user is not equal to expected user. expected: %v received: %v", expectedUser, user)
			}
		})
	}
}

func TestExistsByEmail(t *testing.T) {
	r := NewPGRepository(internaltest.GetPgPool(t))

	type args struct {
		ctx   context.Context
		email string
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "checking for existing user should return true",
			args: args{
				ctx:   context.TODO(),
				email: createRandomUser(t, r).Email,
			},
			want: true,
		},
		{
			name: "checking for non-existing user should return false",
			args: args{
				ctx:   context.TODO(),
				email: gofakeit.Email(),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uExists, err := r.ExistsByEmail(tt.args.ctx, tt.args.email)

			if tt.wantErr != (err != nil) {
				t.Errorf("wantErr expected: %v received: %v", tt.wantErr, err)
			}

			if uExists != tt.want {
				t.Errorf("uExists not equal to tt.wanted expected: %v received: %v", tt.want, uExists)
			}
		})
	}
}
