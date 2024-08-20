-- +goose Up
-- +goose StatementBegin
CREATE COLLATION case_insensitive (provider = icu, locale = 'und-u-ks-level2', deterministic = false);

CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(16) NOT NULL UNIQUE COLLATE case_insensitive,
    email TEXT NOT NULL UNIQUE COLLATE case_insensitive,
    password TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
