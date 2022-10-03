-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS services (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() ,
    name text DEFAULT NULL,
    photo text DEFAULT NULL,
    video text DEFAULT NULL,
    price text DEFAULT NULL,
    description text NOT NULL
);

-- +goose StatementEnd
-- +goose Down