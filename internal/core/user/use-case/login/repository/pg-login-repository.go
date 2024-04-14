package repository

import (
	"context"
	"fmt"

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

func (repo pgRepository) FindPasswordByUsername(ctx context.Context, username model.Username) (DtoPassword, error) {
	res := DtoPassword{}

	sql, args, err := psql.
		Select("id, password").
		From("users").
		Where(
			"username = $1",
			username.String(),
		).
		ToSql()

	if err != nil {
		return res, shared.NewPgRepoErrQueryCreation(err)
	}

	err = repo.db.QueryRow(ctx, sql, args...).Scan(&res.UserID, &res.Value)

	if err == pgx.ErrNoRows {
		return res, fmt.Errorf("password for username %s not found: %w", username.String(), shared.ErrNotFound)
	}

	if err != nil {
		return res, shared.NewPgRepoErrExecution(sql, args, err)
	}

	return res, nil
}

func (repo pgRepository) SaveSession(ctx context.Context, dto DtoCreateSession) (DtoSession, error) {
	sql, args, err := psql.
		Insert("sessions").
		Columns("\"user\"", "creation_date").
		Values(
			dto.UserID,
			dto.CreationDate,
		).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		return DtoSession{}, shared.NewPgRepoErrQueryCreation(err)
	}

	var id int
	err = repo.db.QueryRow(ctx, sql, args...).Scan(&id)

	if err != nil {
		return DtoSession{}, shared.NewPgRepoErrExecution(sql, args, err)
	}

	return DtoSession{ID: id, CreationDate: dto.CreationDate}, nil
}
