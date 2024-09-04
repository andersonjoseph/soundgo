package audio

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/andersonjoseph/soundgo/internal/internaltest"
)

func TestLocalFileRepository_Save(t *testing.T) {
	type args struct {
		ctx context.Context
		i   FileSaveInput
	}

	testReader := bufio.NewReader(bytes.NewBufferString("audio content..."))

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "saving file",
			args: args{
				ctx: context.TODO(),
				i: FileSaveInput{
					ID:   internaltest.GenerateUUID(t),
					file: testReader,
				},
			},
		},
	}

	path, err := os.MkdirTemp("", "soundgo_audios")
	if err != nil {
		t.Fatalf("error while creating temp dir: %v", err)
	}
	defer os.RemoveAll(path)

	r := NewLocalFileRepository(path)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.Save(tt.args.ctx, tt.args.i)
			if !errors.Is(err, tt.err) {
				t.Errorf("Test failed: err expected: %v. received: %v", tt.err, err)
			}

			if tt.err != nil {
				return
			}
		})
	}
}

func TestLocalFileRepository_Get(t *testing.T) {
	testReader := bufio.NewReader(bytes.NewBufferString("audio content..."))

	path, err := os.MkdirTemp("", "soundgo_audios")
	if err != nil {
		t.Fatalf("error while creating temp dir: %v", err)
	}
	defer os.RemoveAll(path)

	r := NewLocalFileRepository(path)

	id := internaltest.GenerateUUID(t)
	err = r.Save(context.TODO(), FileSaveInput{
		ID:   id,
		file: testReader,
	})

	if err != nil {
		t.Fatalf("error while saving audio file: %v", err)
	}

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
			name: "saving file",
			args: args{
				ctx: context.TODO(),
				id:  id,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := r.Get(tt.args.ctx, tt.args.id)
			if !errors.Is(err, tt.err) {
				t.Errorf("Test failed: err expected: %v. received: %v", tt.err, err)
			}

			if tt.err != nil {
				return
			}

			if f.Size != int64(16) {
				t.Errorf("Test failed: f.size expected: %v. received: %v", 16, f.Size)
			}
		})
	}
}

func TestLocalFileRepository_Remove(t *testing.T) {
	testReader := bufio.NewReader(bytes.NewBufferString("audio content..."))

	path, err := os.MkdirTemp("", "soundgo_audios")
	if err != nil {
		t.Fatalf("error while creating temp dir: %v", err)
	}
	defer os.RemoveAll(path)

	r := NewLocalFileRepository(path)

	id := internaltest.GenerateUUID(t)
	err = r.Save(context.TODO(), FileSaveInput{
		ID:   id,
		file: testReader,
	})

	if err != nil {
		t.Fatalf("error while saving audio file: %v", err)
	}

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
			name: "saving file",
			args: args{
				ctx: context.TODO(),
				id:  id,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.Remove(tt.args.ctx, tt.args.id)
			if !errors.Is(err, tt.err) {
				t.Errorf("Test failed: err expected: %v. received: %v", tt.err, err)
			}

			if _, err := os.Stat(filepath.Join(path, id)); !errors.Is(err, os.ErrNotExist) {
				t.Errorf("Test failed: file %v still exists after deletion", id)
			}
		})
	}
}
