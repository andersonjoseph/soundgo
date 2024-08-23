package password

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type PGRepository struct {
	db *pgxpool.Pool
}

func NewPGRepository(db *pgxpool.Pool) PGRepository {
	return PGRepository{
		db: db,
	}
}

func (r PGRepository) SaveResetRequest(ctx context.Context, i SaveResetRequestInput) error {
	query, args, err := psql.
		Insert("password_reset_requests").
		Columns("code", "expires_at", "user_id").
		Values(i.code, i.expiresAt, i.userID).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return handlePgError(err)
	}

	return err
}

func (r PGRepository) GetRequestByUserID(ctx context.Context, userID string) (e RequestReset, err error) {
	query, args, err := psql.
		Select("id", "code", "expires_at", "user_id").
		From("password_reset_requests").
		Where("user_id = ?", userID).
		ToSql()

	if err != nil {
		return e, err
	}

	if err = r.db.QueryRow(ctx, query, args...).Scan(&e.ID, &e.Code, &e.ExpiresAt, &e.UserID); err != nil {
		return e, handlePgError(err)
	}

	return e, err
}

func handlePgError(err error) error {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) && pgErr.Code == "23503" {
		return fmt.Errorf("%w: %s", shared.ErrNotFound, pgErr.Detail)
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return shared.ErrNotFound
	}

	return err
}
