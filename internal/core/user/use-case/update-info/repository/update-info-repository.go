package repository

import (
	"context"

	"github.com/andersonjoseph/soundgo/internal/core/user/model"
)

type UpdateInfoParams struct {
	ID       int
	Username model.Username
}

type UpdateInfoRepository interface {
	UpdateInfo(context context.Context, params UpdateInfoParams) error
}
