package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type pgRepository struct {
	db *pgxpool.Pool
}

func NewPGRepository(db *pgxpool.Pool) pgRepository {
	return pgRepository{db: db}
}

func (repo pgRepository) UpdateInfo(ctx context.Context, params UpdateInfoParams) error {
	queryBuilder := psql.
		Update("\"users\"")

	if len(params.Username.String()) != 0 {
		queryBuilder = queryBuilder.Set("username", params.Username.String())
	}

	sql, args, err := queryBuilder.Where(
		squirrel.Eq{"id": params.ID},
	).
		ToSql()

	if err != nil {
		return shared.NewPgRepoErrQueryCreation(err)
	}

	tag, err := repo.db.Exec(ctx, sql, args...)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return fmt.Errorf("%s: %w", pgErr.Message, shared.ErrRecordAlreadyExists)
		}

		return shared.NewPgRepoErrExecution(sql, args, err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("user with id: %d: %w", params.ID, shared.ErrNotFound)
	}

	return nil
}
