package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/andersonjoseph/soundgo/internal/core/user/model"
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type pgRepository struct {
	db *pgxpool.Pool
}

func NewPGRepository(db *pgxpool.Pool) pgRepository {
	return pgRepository{db: db}
}

func (repo pgRepository) FindPasswordByID(ctx context.Context, id int) (string, error) {
	sql, args, err := psql.
		Select("password").
		From("users").
		Where(
			"id = $1",
			id,
		).
		ToSql()

	if err != nil {
		return "", shared.NewPgRepoErrQueryCreation(err)
	}

	var password string

	err = repo.db.QueryRow(ctx, sql, args...).Scan(&password)

	if err != nil {
		return "", shared.NewPgRepoErrExecution(sql, args, err)
	}

	return password, nil
}

func (repo pgRepository) UpdatePassword(ctx context.Context, id int, password model.Password) error {
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		return shared.NewPgRepoErrTransaction("begin", err)
	}
	defer tx.Rollback(ctx)

	sql, args, err := psql.
		Update("\"users\"").
		Set("password", password.String()).
		Where(
			"id = $2",
			id,
		).
		ToSql()

	if err != nil {
		return shared.NewPgRepoErrQueryCreation(err)
	}

	tag, err := tx.Exec(ctx, sql, args...)

	if err != nil {
		return shared.NewPgRepoErrExecution(sql, args, err)
	}

	if tag.RowsAffected() == 0 {
		return nil
	}

	err = deleteSessions(ctx, tx, id)

	if err != nil {
		return err
	}

	err = tx.Commit(ctx)

	if err != nil {
		return shared.NewPgRepoErrTransaction("commit", err)
	}

	return nil
}

func deleteSessions(ctx context.Context, tx pgx.Tx, userID int) error {
	sql, args, err := psql.
		Delete("sessions").
		Where(
			"\"user\" = $1",
			userID,
		).
		ToSql()

	if err != nil {
		return shared.NewPgRepoErrQueryCreation(err)
	}

	_, err = tx.Exec(ctx, sql, args...)

	if err != nil {
		return shared.NewPgRepoErrExecution(sql, args, err)
	}

	return err
}
