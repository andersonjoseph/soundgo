package audio

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/internaltest"
	"github.com/brianvoe/gofakeit/v7"
)

type testPlayCountRepo struct {
	ch        chan struct{}
	returnErr bool
}

func (r testPlayCountRepo) SavePlayCount(ctx context.Context, id string, count uint64) (uint64, error) {
	if r.returnErr {
		return 0, fmt.Errorf("oops, an error occured")
	}

	r.ch <- struct{}{}
	return count, nil
}

func TestMemoryPlayCountHandler_Add(t *testing.T) {
	testCh := make(chan struct{})

	type args struct {
		ctx      context.Context
		playerID string
		audio    Entity
	}

	tests := []struct {
		name      string
		args      args
		repo      playCountRepository
		size      uint64
		interval  time.Duration
		err       error
		wantChErr bool
		wantErr   bool
	}{
		{
			name:     "add a play",
			repo:     testPlayCountRepo{ch: testCh},
			size:     2,
			interval: time.Second * 1,
			args: args{
				ctx:      context.TODO(),
				playerID: internaltest.GenerateUUID(t),
				audio: Entity{
					ID:        internaltest.GenerateUUID(t),
					Title:     gofakeit.BookTitle(),
					Status:    api.AudioStatusPublished,
					UserID:    internaltest.GenerateUUID(t),
					CreatedAt: time.Now(),
				},
			},
		},
		{
			name:      "error in repository (async)",
			wantChErr: true,
			repo:      testPlayCountRepo{ch: testCh, returnErr: true},
			size:      2,
			interval:  time.Second * 1,
			args: args{
				ctx:      context.TODO(),
				playerID: internaltest.GenerateUUID(t),
				audio: Entity{
					ID:        internaltest.GenerateUUID(t),
					Title:     gofakeit.BookTitle(),
					Status:    api.AudioStatusPublished,
					UserID:    internaltest.GenerateUUID(t),
					CreatedAt: time.Now(),
				},
			},
		},
	}

	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countHandler := NewMemoryPlayCountHandler(context.TODO(), tt.size, tt.repo, tt.interval, logger)

			err := countHandler.Add(tt.args.ctx, tt.args.playerID, tt.args.audio)
			if tt.wantErr != (err != nil) {
				t.Errorf("Test failed: received unexpected error: %v", err)
			}

			if tt.wantChErr {
				select {
				case <-countHandler.errCh:
					break
				case <-time.After(time.Second * 5):
					t.Error("Test failed: expected async error timed out")
				}
			} else {
				select {
				case <-testCh:
					break
				case <-time.After(time.Second * 5):
					t.Error("Test failed: save to repository timed out")
				}
			}
		})
	}
}
