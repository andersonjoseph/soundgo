-- +goose Up
-- +goose StatementBegin

CREATE COLLATION case_insensitive (LOCALE = 'en_US', PROVIDER = 'icu', DETERMINISTIC = false);

CREATE TABLE users (
    id                  serial PRIMARY KEY,
    email               VARCHAR(254) NOT NULL UNIQUE COLLATE case_insensitive,
    username            VARCHAR(32) NOT NULL UNIQUE COLLATE case_insensitive,
    password            TEXT NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users
DROP COLLATION case_insensitive
-- +goose StatementEnd
