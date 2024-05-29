package updateuserinfo

import (
	"context"

	"github.com/andersonjoseph/soundgo/internal/core/user/model"
	"github.com/andersonjoseph/soundgo/internal/core/user/use-case/update-info/repository"
)

type useCase struct {
	repository repository.UpdateInfoRepository
}

func New(repo repository.UpdateInfoRepository) useCase {
	return useCase{
		repository: repo,
	}
}

type UpdateInfoParams struct {
	ID       int
	Username string
}

func (s useCase) UpdateInfo(ctx context.Context, params UpdateInfoParams) error {
	var err error
	var u model.Username

	if len(params.Username) > 0 {
		u, err = model.NewUsername(params.Username)

		if err != nil {
			return err
		}
	}

	return s.repository.UpdateInfo(ctx, repository.UpdateInfoParams{
		ID:       params.ID,
		Username: u,
	})
}
