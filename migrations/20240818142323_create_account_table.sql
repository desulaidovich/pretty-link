-- +goose Up
-- +goose StatementBegin
CREATE TABLE account (
    id SERIAL PRIMARY KEY,
    email CHARACTER VARYING(255) UNIQUE,
    password BYTEA NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CHECK ((email != ''))
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS account CASCADE;
-- +goose StatementEnd
