package audio

import (
	"context"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/reqcontext"
)

func (h Handler) GetUserAudios(ctx context.Context, params api.GetUserAudiosParams) (api.GetUserAudiosRes, error) {
	currUser, err := reqcontext.CurrentUserID.Get(ctx)

	excludeHidden := err != nil || currUser != params.ID
	audios, err := h.repository.GetByUser(ctx, params.ID, params.XPaginationAfter.Value, uint64(params.XPaginationLimit.Or(20)), excludeHidden)

	var paginationNext string
	if len(audios) > 0 {
		paginationNext = audios[len(audios)-1].ID
	}

	audioResponses := make([]api.Audio, len(audios))
	for i := range audioResponses {
		audioResponses[i] = api.Audio{
			ID:          audios[i].ID,
			Title:       audios[i].Title,
			Description: api.NewOptString(audios[i].Description),
			CreatedAt:   audios[i].CreatedAt,
			User:        audios[i].UserID,
			Status:      audios[i].Status,
			PlayCount:   audios[i].Playcount,
		}
	}

	return &api.GetUserAudiosOKHeaders{
		XPaginationNext: api.NewOptString(paginationNext),
		Response:        audioResponses,
	}, nil
}
