package audio

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/andersonjoseph/soundgo/internal/shared"
)

type LocalFileRepository struct {
	basePath string
}

func NewLocalFileRepository(basePath string) LocalFileRepository {
	return LocalFileRepository{
		basePath: basePath,
	}
}

func (r LocalFileRepository) Save(ctx context.Context, i FileSaveInput) error {
	path := filepath.Join(r.basePath, i.ID)

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error while creating file: %w", err)
	}

	_, err = io.Copy(f, i.file)
	if err != nil {
		return fmt.Errorf("error while copying file: %w", err)
	}

	return nil
}

func (r LocalFileRepository) Get(ctx context.Context, ID string) (File, error) {
	path := filepath.Join(r.basePath, ID)

	file, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return File{}, shared.ErrNotFound
	}
	if err != nil {
		return File{}, fmt.Errorf("error while opening file: %w", err)
	}

	info, err := file.Stat()
	if err != nil {
		return File{}, fmt.Errorf("error while getting file info: %w", err)
	}

	return File{
		File: file,
		Size: info.Size(),
	}, nil
}

func (r LocalFileRepository) Remove(ctx context.Context, ID string) error {
	path := filepath.Join(r.basePath, ID)

	err := os.Remove(path)
	if errors.Is(err, os.ErrNotExist) {
		return shared.ErrNotFound
	}
	if err != nil {
		return fmt.Errorf("error while deleting audio file: %w", err)
	}

	return nil
}
