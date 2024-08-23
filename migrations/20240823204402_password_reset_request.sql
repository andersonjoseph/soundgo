-- +goose Up
-- +goose StatementBegin
CREATE TABLE password_reset_requests (
    id SERIAL PRIMARY KEY,
    code VARCHAR(6) NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    expires_at timestamptz NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP password_reset_request
-- +goose StatementEnd
