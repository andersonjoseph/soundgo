package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
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

func (repo pgRepository) Save(ctx context.Context, dto DtoSaveUser) (DtoUser, error) {
	if err := repo.errIfUserExists(ctx, dto); err != nil {
		return DtoUser{}, err
	}

	sql, args, err := psql.
		Insert("users").
		Columns("email", "username", "password").
		Values(
			dto.Email.String(),
			dto.Username.String(),
			dto.Password.String(),
		).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		return DtoUser{}, shared.NewPgRepoErrQueryCreation(err)
	}

	tx, err := repo.db.Begin(ctx)
	if err != nil {
		return DtoUser{}, shared.NewPgRepoErrTransaction("begin", err)
	}

	defer tx.Rollback(ctx)

	var id int
	err = tx.QueryRow(ctx, sql, args...).Scan(&id)

	if err != nil {
		return DtoUser{}, shared.NewPgRepoErrExecution(sql, args, err)
	}

	u, err := findById(ctx, tx, id)

	if err != nil {
		return DtoUser{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return DtoUser{}, shared.NewPgRepoErrTransaction("commit", err)
	}

	return u, nil
}

func (repo pgRepository) userExistsBy(ctx context.Context, col string, pred string) (bool, error) {
	sql, args, err := psql.
		Select("count(id) > 0").
		From("users").
		Where(
			squirrel.Eq{col: pred},
		).
		ToSql()

	if err != nil {
		return false, shared.NewPgRepoErrQueryCreation(err)
	}

	var exists bool
	if err := repo.db.QueryRow(ctx, sql, args...).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (repo pgRepository) errIfUserExists(ctx context.Context, dto DtoSaveUser) error {
	if userExists, err := repo.userExistsBy(ctx, "username", dto.Username.String()); err != nil {
		return err
	} else if userExists {
		return shared.NewPgRepoErrExistingRecord("username", dto.Username.String())
	}

	if userExists, err := repo.userExistsBy(ctx, "email", dto.Email.String()); err != nil {
		return err
	} else if userExists {
		return shared.NewPgRepoErrExistingRecord("email", dto.Email.String())
	}

	return nil
}

func findById(ctx context.Context, tx pgx.Tx, id int) (DtoUser, error) {
	sql, args, err := psql.
		Select("id", "email", "username", "password").
		From("users").
		Where(
			"id = $1",
			id,
		).
		ToSql()

	if err != nil {
		return DtoUser{}, shared.NewPgRepoErrQueryCreation(err)
	}

	u := DtoUser{}

	err = tx.QueryRow(ctx, sql, args...).Scan(&u.Id, &u.Email, &u.Username, &u.Password)

	if err != nil {
		return DtoUser{}, shared.NewPgRepoErrExecution(sql, args, err)
	}

	return u, nil
}
