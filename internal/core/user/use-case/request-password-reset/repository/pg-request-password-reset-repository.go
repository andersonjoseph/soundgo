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

func (repo pgRepository) FindByUsernameAndEmail(ctx context.Context, username model.Username, email model.Email) (FindByUserDTO, error) {
	res := FindByUserDTO{}

	sql, args, err := psql.
		Select("id, email, username").
		From("users").
		Where(
			"username = $1",
			username.String(),
		).
		Where(
			"email = $2",
			email.String(),
		).
		ToSql()

	if err != nil {
		return res, shared.NewPgRepoErrQueryCreation(err)
	}

	err = repo.db.QueryRow(ctx, sql, args...).Scan(&res.ID, &res.Email, &res.Username)

	if err == pgx.ErrNoRows {
		return res, shared.ErrNotFound
	}

	if err != nil {
		return FindByUserDTO{}, shared.NewPgRepoErrExecution(sql, args, err)
	}

	return res, nil
}
