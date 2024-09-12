package audio

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/andersonjoseph/soundgo/internal/api"
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
	queryBuilder := psql.Update("audios")

	updateIfNotZero(&queryBuilder, "title", i.Title)
	updateIfNotZero(&queryBuilder, "description", i.Description)
	updateIfNotZero(&queryBuilder, "status", i.Status)
	updateIfNotZero(&queryBuilder, "play_count", i.PlayCount)

	query, args, err := queryBuilder.
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

func (r PGRepository) SavePlayCount(ctx context.Context, ID string, count uint64) (uint64, error) {
	query, args, err := psql.Update("audios").
		Set("play_count", squirrel.Expr("play_count + ?", count)).
		Suffix("RETURNING play_count").
		Where("ID = ?", ID).
		ToSql()

	if err != nil {
		return 0, err
	}

	if err = r.db.QueryRow(ctx, query, args...).Scan(&count); err != nil {
		return 0, handlePgError(err)
	}

	return count, err
}

func (r PGRepository) GetByUser(ctx context.Context, userID string, after string, limit uint64, excludeHidden bool) ([]Entity, error) {
	if limit == 0 {
		limit = 20
	}

	builder := psql.
		Select("id", "title", "description", "play_count", "status", "created_at", "user_id").
		From("audios").
		Where("user_id = ?", userID).
		OrderBy("id ASC").
		Limit(limit)

	if after != "" {
		builder = builder.Where("id > ?", after)
	}
	if excludeHidden {
		builder = builder.Where("status != 'hidden'")
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	var entities []Entity
	for rows.Next() {
		var e Entity
		err := rows.Scan(
			&e.ID,
			&e.Title,
			&e.Description,
			&e.Playcount,
			&e.Status,
			&e.CreatedAt,
			&e.UserID,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		entities = append(entities, e)
	}

	if err = rows.Err(); err != nil {
		return nil, handlePgError(err)
	}

	return entities, nil
}

func updateIfNotZero(builder *squirrel.UpdateBuilder, colName string, val any) {
	switch t := val.(type) {
	case string, api.UpdateAudioInputStatus:
		if t != "" {
			*builder = builder.Set(colName, val)
		}
	case uint64:
		if t != 0 {
			*builder = builder.Set(colName, val)
		}
	default:
		panic(fmt.Sprintf("type: %v of value: %v cannot be checked for a zero-value", t, val))
	}
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
