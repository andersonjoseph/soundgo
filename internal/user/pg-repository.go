package user

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
		Insert("users").
		Columns("id", "username", "email", "password").
		Values(i.ID, i.Username, i.Email, i.Password).
		Suffix("RETURNING id, username, email, password, created_at").
		ToSql()

	if err != nil {
		return e, err
	}

	if err = r.db.QueryRow(ctx, query, args...).Scan(&e.ID, &e.Username, &e.Email, &e.Password, &e.CreatedAt); err != nil {
		return e, handlePgError(err)
	}

	return e, err
}

func (r PGRepository) Update(ctx context.Context, ID string, i UpdateInput) (e Entity, err error) {
	query, args, err := psql.
		Update("users").
		Set("username", i.Username).
		Suffix("RETURNING id, username, email, password, created_at").
		Where("ID = ?", ID).
		ToSql()

	if err != nil {
		return e, err
	}

	if err = r.db.QueryRow(ctx, query, args...).Scan(&e.ID, &e.Username, &e.Email, &e.Password, &e.CreatedAt); err != nil {
		return e, handlePgError(err)
	}

	return e, err
}

func (r PGRepository) GetByUsername(ctx context.Context, username string) (e Entity, err error) {
	return r.getBy(ctx, "username", username)
}

func (r PGRepository) GetByEmail(ctx context.Context, email string) (e Entity, err error) {
	return r.getBy(ctx, "email", email)
}

func (r PGRepository) ExistsByEmail(ctx context.Context, email string) (b bool, err error) {
	query, args, err := psql.
		Select("count(id) > 0").
		From("users").
		Where("email = ?", email).
		ToSql()

	if err != nil {
		return b, err
	}

	if err = r.db.QueryRow(ctx, query, args...).Scan(&b); err != nil {
		return b, handlePgError(err)
	}

	return b, err
}

func (r PGRepository) getBy(ctx context.Context, column, value string) (e Entity, err error) {
	query, args, err := psql.
		Select("id", "username", "email", "password", "created_at").
		From("users").
		Where(fmt.Sprintf("%s = ?", column), value).
		ToSql()

	if err != nil {
		return e, err
	}

	if err = r.db.QueryRow(ctx, query, args...).Scan(&e.ID, &e.Username, &e.Email, &e.Password, &e.CreatedAt); err != nil {
		return e, handlePgError(err)
	}

	return e, err
}

func handlePgError(err error) error {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return fmt.Errorf("%w: %s", shared.ErrAlreadyExists, pgErr.Detail)
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return shared.ErrNotFound
	}

	return err
}
