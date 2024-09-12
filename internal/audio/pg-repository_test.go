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

func TestDelete(t *testing.T) {
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
			name: "deleting audio",
			args: args{
				ctx: context.TODO(),
				id:  audio.ID,
			},
		},
		{
			name: "deleting non existing audio",
			args: args{
				ctx: context.TODO(),
				id:  internaltest.GenerateUUID(t),
			},
			err: shared.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.Delete(tt.args.ctx, tt.args.id)
			if !errors.Is(err, tt.err) {
				t.Errorf("Test failed: err expected: %v. received: %v", tt.err, err)
			}

			if tt.err != nil {
				return
			}
		})
	}

}

func TestUpdate(t *testing.T) {
	pool := internaltest.GetPgPool(t)
	userRepo := user.NewPGRepository(pool)
	repo := NewPGRepository(pool)

	audio := createRandomAudio(t, repo, userRepo)

	type args struct {
		ctx context.Context
		id  string
		i   UpdateInput
	}

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "updating audio",
			args: args{
				ctx: context.TODO(),
				id:  audio.ID,
				i: UpdateInput{
					Title:       gofakeit.BookTitle(),
					Description: gofakeit.Name(),
					Status:      api.UpdateAudioInputStatus(audio.Status),
				},
			},
		},
		{
			name: "updating non existing audio",
			args: args{
				ctx: context.TODO(),
				id:  internaltest.GenerateUUID(t),
				i: UpdateInput{
					Title:       gofakeit.BookTitle(),
					Description: gofakeit.Name(),
					Status:      api.UpdateAudioInputStatus(audio.Status),
				},
			},
			err: shared.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := repo.Update(tt.args.ctx, tt.args.id, tt.args.i)
			if !errors.Is(err, tt.err) {
				t.Errorf("Test failed: err expected: %v. received: %v", tt.err, err)
			}

			if tt.err != nil {
				return
			}

			if e.ID != tt.args.id {
				t.Errorf("Test failed: ID expected: %v. received: %v", tt.args.id, e.ID)
			}

			if e.Title != tt.args.i.Title {
				t.Errorf("Test failed: title expected: %v. received: %v", tt.args.i.Title, e.Title)
			}

			if e.Description != tt.args.i.Description {
				t.Errorf("Test failed: Description expected: %v. received: %v", tt.args.i.Description, e.Description)
			}

			if e.UserID != audio.UserID {
				t.Errorf("Test failed: UserID expected: %v. received: %v", audio.UserID, e.UserID)
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

func TestSavePlayCount(t *testing.T) {
	pool := internaltest.GetPgPool(t)
	userRepo := user.NewPGRepository(pool)
	repo := NewPGRepository(pool)

	type args struct {
		ctx   context.Context
		id    string
		count uint64
	}

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "incrementing play count",
			args: args{
				ctx:   context.TODO(),
				id:    createRandomAudio(t, repo, userRepo).ID,
				count: 1,
			},
		},
		{
			name: "incrementing play count by 10",
			args: args{
				ctx:   context.TODO(),
				id:    createRandomAudio(t, repo, userRepo).ID,
				count: 10,
			},
		},
		{
			name: "incrementing play count for non existing audio",
			args: args{
				ctx:   context.TODO(),
				id:    internaltest.GenerateUUID(t),
				count: 1,
			},
			err: shared.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count, err := repo.SavePlayCount(tt.args.ctx, tt.args.id, tt.args.count)
			if !errors.Is(err, tt.err) {
				t.Errorf("Test failed: err expected: %v. received: %v", tt.err, err)
			}

			if tt.err != nil {
				return
			}

			if count != tt.args.count {
				t.Errorf("Test failed: count expected: %v. received: %v", tt.args.count, count)
			}
		})
	}
}

func TestGetByUser(t *testing.T) {
	pool := internaltest.GetPgPool(t)
	userRepo := user.NewPGRepository(pool)
	repo := NewPGRepository(pool)

	type args struct {
		ctx    context.Context
		userID string
		after  string
		limit  uint64
	}

	tests := []struct {
		name  string
		args  args
		count int
		err   error
	}{
		{
			name:  "get audios by user",
			count: 5,
			args: args{
				ctx:   context.TODO(),
				after: "",
				limit: 0,
			},
		},
		{
			name:  "get audios by user with limit",
			count: 5,
			args: args{
				ctx:   context.TODO(),
				after: "",
				limit: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := createRandomUser(t, userRepo)
			tt.args.userID = user.ID

			for i := 0; i < tt.count; i++ {
				createAudio(t, repo, SaveInput{
					ID:          internaltest.GenerateUUID(t),
					Title:       gofakeit.BookTitle(),
					Description: gofakeit.Name(),
					UserID:      user.ID,
					Status:      api.AudioInputMultipartStatusPublished,
				})
			}

			audios, err := repo.GetByUser(tt.args.ctx, tt.args.userID, tt.args.after, tt.args.limit, false)
			if !errors.Is(err, tt.err) {
				t.Errorf("Test failed: err expected: %v. received: %v", tt.err, err)
			}
			if tt.err != nil {
				return
			}

			if len(audios) != tt.count && len(audios) != int(tt.args.limit) {
				t.Errorf("Test failed: audios length expected: %v. received: %v", tt.count, len(audios))
			}
		})
	}
}

func TestGetByUserHidden(t *testing.T) {
	pool := internaltest.GetPgPool(t)
	userRepo := user.NewPGRepository(pool)
	repo := NewPGRepository(pool)

	type args struct {
		ctx    context.Context
		userID string
		after  string
		limit  uint64
	}

	tests := []struct {
		name  string
		args  args
		count int
		err   error
	}{
		{
			name:  "get audios by user excluding hidden",
			count: 5,
			args: args{
				ctx:   context.TODO(),
				after: "",
				limit: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := createRandomUser(t, userRepo)
			tt.args.userID = user.ID

			for i := 0; i < tt.count; i++ {
				createAudio(t, repo, SaveInput{
					ID:          internaltest.GenerateUUID(t),
					Title:       gofakeit.BookTitle(),
					Description: gofakeit.Name(),
					UserID:      user.ID,
					Status:      api.AudioInputMultipartStatusHidden,
				})
			}

			audios, err := repo.GetByUser(tt.args.ctx, tt.args.userID, tt.args.after, tt.args.limit, true)
			if !errors.Is(err, tt.err) {
				t.Errorf("Test failed: err expected: %v. received: %v", tt.err, err)
			}
			if tt.err != nil {
				return
			}

			if len(audios) != 0 {
				t.Errorf("Test failed: audios length expected: %v. received: %v", 0, len(audios))
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

func createAudio(t *testing.T, repo PGRepository, i SaveInput) Entity {
	t.Helper()

	a, err := repo.Save(context.TODO(), i)

	if err != nil {
		t.Fatalf("Test failed: error occured while creating test user. received: %v", err)
	}

	return a
}

func createRandomAudio(t *testing.T, repo PGRepository, userRepo user.PGRepository) Entity {
	t.Helper()

	return createAudio(t, repo, SaveInput{
		ID:          internaltest.GenerateUUID(t),
		Title:       gofakeit.BookTitle(),
		Description: gofakeit.Name(),
		UserID:      createRandomUser(t, userRepo).ID,
		Status:      api.AudioInputMultipartStatusPublished,
	})
}
