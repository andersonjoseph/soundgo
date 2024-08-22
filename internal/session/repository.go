package session

import "context"

type SaveInput struct {
	ID     string
	Token  string
	UserID string
}

type Repository interface {
	Save(ctx context.Context, i SaveInput) (Entity, error)
	Delete(ctx context.Context, ID string) error
}
