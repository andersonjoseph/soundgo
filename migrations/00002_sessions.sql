-- +goose Up
-- +goose StatementBegin
CREATE TABLE sessions (
    id              serial PRIMARY KEY,
    "user"          integer NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    creation_date   timestamp with time zone
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sessions;
-- +goose StatementEnd
