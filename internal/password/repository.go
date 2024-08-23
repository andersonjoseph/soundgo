package password

import (
	"context"
	"time"
)

type SaveResetRequestInput struct {
	userID    string
	code      string
	expiresAt time.Time
}

type Repository interface {
	SaveResetRequest(ctx context.Context, i SaveResetRequestInput) error
	GetRequestByUserID(ctx context.Context, userID string) (RequestReset, error)
}
