package audio

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

func (r PGRepository) Save(ctx context.Context, i SaveInput) (e Entity, err error) {
	query, args, err := psql.
		Insert("audios").
		Columns("id", "title", "description", "status", "user_id").
		Values(i.ID, i.Title, i.Description, i.Status, i.UserID).
		Suffix("RETURNING id, title, description, play_count, status, created_at, user_id").
		ToSql()

	if err != nil {
		return e, err
	}

	if err = r.db.QueryRow(ctx, query, args...).Scan(&e.ID, &e.Title, &e.Description, &e.Playcount, &e.Status, &e.CreatedAt, &e.UserID); err != nil {
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
