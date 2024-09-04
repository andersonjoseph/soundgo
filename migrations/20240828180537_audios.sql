-- +goose Up
-- +goose StatementBegin

CREATE TYPE audio_status AS ENUM ('published', 'pending', 'hidden');

CREATE TABLE audios (
    id UUID PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(5000),
    play_count BIGINT NOT NULL DEFAULT 0,
    status audio_status NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE audios;
-- +goose StatementEnd
