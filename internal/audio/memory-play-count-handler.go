package audio

import (
	"context"
	"strings"
	"sync"
	"time"
)

type playCountRepository interface {
	SavePlayCount(ctx context.Context, id string, count uint64) (uint64, error)
}

type memoryPlayCountHandler struct {
	set   *safeSet
	repo  playCountRepository
	errCh chan error
}

func NewMemoryPlayCountHandler(ctx context.Context, size uint64, repo playCountRepository, saveInterval time.Duration) memoryPlayCountHandler {
	set := newSafeSet(size)
	errCh := make(chan error)

	go func() {
		t := time.NewTicker(saveInterval)
		defer t.Stop()

		for {
			select {
			case <-ctx.Done():
				return

			case <-t.C:
				err := processCounts(ctx, &set, repo)
				if err != nil {
					errCh <- err
				}
			}
		}
	}()

	return memoryPlayCountHandler{
		set:   &set,
		repo:  repo,
		errCh: errCh,
	}
}

func (h memoryPlayCountHandler) Add(ctx context.Context, playerID string, audio Entity) error {
	hasSpace := h.set.add(audio.ID + ":" + playerID)

	if !hasSpace {
		err := processCounts(ctx, h.set, h.repo)
		if err != nil {
			return err
		}
		return h.Add(ctx, playerID, audio)
	}

	return nil
}

func processCounts(ctx context.Context, set *safeSet, repo playCountRepository) error {
	if set.size() == 0 {
		return nil
	}

	playCounts := make(map[string]uint64)

	set.process(func(s string) {
		id := strings.Split(s, ":")[0]
		playCounts[id]++
	})

	for id, count := range playCounts {
		if _, err := repo.SavePlayCount(ctx, id, count); err != nil {
			return err
		}
		delete(playCounts, id)
	}

	playCounts = nil
	return nil
}

type safeSet struct {
	set     map[string]struct{}
	mu      sync.RWMutex
	maxSize uint64
}

func newSafeSet(size uint64) safeSet {
	return safeSet{
		mu:      sync.RWMutex{},
		set:     make(map[string]struct{}, size),
		maxSize: size,
	}
}

func (s *safeSet) add(id string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.maxSize == uint64(len(s.set)) {
		return false
	}

	s.set[id] = struct{}{}
	return true
}

func (s *safeSet) process(fn func(string)) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for id := range s.set {
		fn(id)
	}

	s.set = make(map[string]struct{}, s.maxSize)
}

func (s *safeSet) size() int {
	return len(s.set)
}
