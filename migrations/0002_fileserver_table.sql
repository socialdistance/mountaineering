-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS files (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() ,
    name text DEFAULT NULL,
    path text DEFAULT NULL,
    description text NOT NULL
);

-- +goose StatementEnd
-- +goose Down