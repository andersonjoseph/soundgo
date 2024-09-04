package audio

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/audio/audiorange"
)

type audioFile struct {
	File *os.File
	Info fs.FileInfo
}

// GET /audios/{id}
func (h Handler) GetAudio(ctx context.Context, params api.GetAudioParams) (api.GetAudioRes, error) {
	audioFile, err := getAudioFile()
	if err != nil {
		return nil, err
	}

	var res api.GetAudioRes
	h.logger.Info("getting audio", slog.Group(
		"input",
		"ID",
		params.ID,
		"range",
		params.Range.Value,
	))

	if params.Range.IsSet() {
		p := audiorange.Parser{}
		parsedRanges, err := p.Parse(params.Range.Value)
		if err != nil {
			return &api.GetAudioRequestedRangeNotSatisfiable{
				Error: api.NewOptString(err.Error()),
			}, nil
		}

		partialFile, err := getPartialFile(audioFile, parsedRanges[0])

		res = &api.GetAudioPartialContentHeaders{
			ContentRange: api.NewOptString(getContentRangeString(parsedRanges[0], audioFile.Info.Size())),
			Response: api.GetAudioPartialContent{
				Data: partialFile,
			},
		}
	} else {
		res = &api.GetAudioOKHeaders{
			AcceptRanges: api.NewOptString("bytes"),
			Response: api.GetAudioOK{
				Data: audioFile.File,
			},
		}
	}

	return res, nil
}

func getPartialFile(f audioFile, fRange audiorange.Range) (io.Reader, error) {
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
		limit = f.Info.Size() - 1
	}

	return io.LimitReader(f.File, limit), nil
}

func getAudioFile() (audioFile, error) {
	path, err := filepath.Abs("audios/audio.m4a")
	if err != nil {
		return audioFile{}, fmt.Errorf("error while building path: %w", err)
	}

	file, err := os.Open(path)
	if err != nil {
		return audioFile{}, fmt.Errorf("error while opening file: %w", err)
	}

	info, err := file.Stat()
	if err != nil {
		return audioFile{}, fmt.Errorf("error while getting file info: %w", err)
	}

	return audioFile{
		File: file,
		Info: info,
	}, nil
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
