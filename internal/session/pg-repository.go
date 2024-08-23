package session

import (
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
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
		Insert("sessions").
		Columns("id", "token", "user_id").
		Values(i.ID, i.Token, i.UserID).
		Suffix("RETURNING id, token").
		ToSql()

	if err != nil {
		return e, err
	}

	if err = r.db.QueryRow(ctx, query, args...).Scan(&e.ID, &e.Token); err != nil {
		return e, handlePgError(err)
	}

	return e, err
}

func (r PGRepository) Delete(ctx context.Context, ID string) (err error) {
	query, args, err := psql.
		Delete("sessions").
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

func (r PGRepository) GetByID(ctx context.Context, ID string) (e Entity, err error) {
	query, args, err := psql.
		Select("id", "token", "user_id").
		From("sessions").
		Where("id = ?", ID).
		ToSql()

	if err != nil {
		return e, err
	}

	if err = r.db.QueryRow(ctx, query, args...).Scan(&e.ID, &e.Token, &e.UserID); err != nil {
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
