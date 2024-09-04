package audio

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/gabriel-vasile/mimetype"
)

// POST /audios
func (h Handler) CreateAudio(ctx context.Context, req *api.AudioInputMultipart) (res api.CreateAudioRes, err error) {
	ID, err := shared.GenerateUUID()
	if err != nil {
		return nil, err
	}

	file := bufio.NewReader(req.File.File)
	err = h.saveFile(ctx, ID, file)

	if errors.Is(err, shared.ErrBadInput) {
		return &api.CreateAudioUnsupportedMediaType{Error: api.NewOptString(err.Error())}, nil
	}
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			err = errors.Join(err, h.fileRepository.Remove(ctx, ID))
		}
	}()

	userID, err := h.contextRequestHandler.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	e, err := h.repository.Save(ctx, SaveInput{
		ID:          ID,
		Title:       req.Title,
		Description: req.Description.Value,
		UserID:      userID,
		Status:      api.AudioStatusPublished,
	})

	if err != nil {
		return nil, err
	}

	return &api.Audio{
		ID:          e.ID,
		Title:       e.Title,
		Description: api.NewOptString(e.Description),
		CreatedAt:   e.CreatedAt,
		PlayCount:   e.Playcount,
		User:        e.UserID,
		Status:      e.Status,
	}, nil
}

func (h Handler) saveFile(ctx context.Context, ID string, file *bufio.Reader) error {
	ok, err := isFileValid(file)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("%w: file must be a valid audio file", shared.ErrBadInput)
	}

	err = h.fileRepository.Save(ctx, FileSaveInput{
		ID:   ID,
		file: file,
	})
	if err != nil {
		return err
	}

	return nil
}

func isFileValid(r *bufio.Reader) (bool, error) {
	buf, err := r.Peek(512)
	if err != nil {
		return false, fmt.Errorf("error while reading file header: %w", err)
	}

	t := mimetype.Detect(buf)
	if !strings.HasPrefix(t.String(), "audio") {
		return false, nil
	}

	return true, nil
}
