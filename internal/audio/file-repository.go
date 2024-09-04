package audio

import (
	"context"
	"io"
)

type FileSaveInput struct {
	ID   string
	file io.Reader
}

type File struct {
	File io.ReadSeeker
	Size int64
}

type FileRepository interface {
	Save(context.Context, FileSaveInput) error
	Get(context.Context, string) (File, error)
	Remove(context.Context, string) error
}
