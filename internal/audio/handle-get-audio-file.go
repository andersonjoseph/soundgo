package audio

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/audio/audiorange"
	"github.com/andersonjoseph/soundgo/internal/shared"
)

// GET /audios/{id}/file
func (h Handler) GetAudioFile(ctx context.Context, params api.GetAudioFileParams) (api.GetAudioFileRes, error) {
	e, err := h.repository.Get(ctx, params.ID)
	if errors.Is(err, shared.ErrNotFound) {
		return &api.GetAudioFileNotFound{}, nil
	}
	if err != nil {
		return nil, err
	}

	if e.Status == api.AudioStatusHidden {
		currentUserID, err := h.contextRequestHandler.GetUserID(ctx)
		if err != nil || currentUserID != e.UserID {
			return &api.GetAudioFileForbidden{}, nil
		}
	}

	audioFile, err := h.fileRepository.Get(ctx, params.ID)
	if errors.Is(err, shared.ErrNotFound) {
		return &api.GetAudioFileNotFound{}, nil
	}
	if err != nil {
		return nil, err
	}

	var res api.GetAudioFileRes

	if params.Range.IsSet() {
		p := audiorange.Parser{}
		parsedRanges, err := p.Parse(params.Range.Value)
		if err != nil {
			return &api.GetAudioFileRequestedRangeNotSatisfiable{
				Error: api.NewOptString(err.Error()),
			}, nil
		}

		partialFile, err := getPartialFile(audioFile, parsedRanges[0])

		res = &api.GetAudioFilePartialContentHeaders{
			ContentRange: api.NewOptString(getContentRangeString(parsedRanges[0], audioFile.Size)),
			Response: api.GetAudioFilePartialContent{
				Data: partialFile,
			},
		}
	} else {
		res = &api.GetAudioFileOKHeaders{
			AcceptRanges: api.NewOptString("bytes"),
			Response: api.GetAudioFileOK{
				Data: audioFile.File,
			},
		}
	}

	//h.playcountHandler.Add(userID, params.ID)

	return res, nil
}

func getPartialFile(f File, fRange audiorange.Range) (io.Reader, error) {
	var seekPosition int
	var offset int64

	if fRange.Start == -1 {
		seekPosition = io.SeekEnd
		offset = fRange.End * -1
	} else {
		seekPosition = io.SeekStart
		offset = fRange.Start
	}

	if _, err := f.File.Seek(offset, seekPosition); err != nil {
		return nil, err
	}

	var limit = fRange.End
	if limit == -1 {
		limit = f.Size - 1
	}

	return io.LimitReader(f.File, limit), nil
}

func getContentRangeString(r audiorange.Range, size int64) string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprint("bytes "))
	if r.Start != -1 {
		b.WriteString(fmt.Sprintf("%d", r.Start))
	}

	b.WriteString("-")

	if r.End == -1 {
		b.WriteString(fmt.Sprintf("%d", size-1))
	} else {
		b.WriteString(fmt.Sprintf("%d", r.End))
	}

	b.WriteString(fmt.Sprintf("/%d", size))

	return b.String()
}
