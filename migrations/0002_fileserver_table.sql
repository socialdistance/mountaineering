-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS files (
    id SERIAL PRIMARY KEY,
    name text DEFAULT NULL,
    path text DEFAULT NULL,
    description text NOT NULL
);

-- +goose StatementEnd
-- +goose Down