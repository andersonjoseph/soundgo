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

func (r PGRepository) Get(ctx context.Context, id string) (e Entity, err error) {
	query, args, err := psql.
		Select("id", "title", "description", "play_count", "status", "created_at", "user_id").
		From("audios").
		Where("id = ?", id).
		ToSql()

	if err != nil {
		return e, err
	}

	err = r.db.QueryRow(ctx, query, args...).Scan(
		&e.ID,
		&e.Title,
		&e.Description,
		&e.Playcount,
		&e.Status,
		&e.CreatedAt,
		&e.UserID,
	)

	if err != nil {
		return e, handlePgError(err)
	}

	return e, nil
}

func (r PGRepository) Delete(ctx context.Context, ID string) (err error) {
	query, args, err := psql.
		Delete("audios").
		Where("id = ?", ID).
		ToSql()

	if err != nil {
		return err
	}

	res, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return shared.ErrNotFound
	}

	return nil
}

func (r PGRepository) Update(ctx context.Context, ID string, i UpdateInput) (e Entity, err error) {
	query, args, err := psql.
		Update("audios").
		Set("title", i.Title).
		Set("description", i.Description).
		Set("status", i.Status).
		Suffix("RETURNING id, title, description, play_count, status, created_at, user_id").
		Where("ID = ?", ID).
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
