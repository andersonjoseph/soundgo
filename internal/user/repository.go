package user

import "context"

type SaveInput struct {
	ID       string
	Username string
	Email    string
	Password string
}

type UpdateInput struct {
	Username string
	Password string
}

type Repository interface {
	Save(ctx context.Context, i SaveInput) (Entity, error)
	Update(ctx context.Context, ID string, i UpdateInput) (Entity, error)
	GetByUsername(ctx context.Context, username string) (Entity, error)
	GetByEmail(ctx context.Context, email string) (Entity, error)

	ExistsByEmail(ctx context.Context, email string) (bool, error)
}
