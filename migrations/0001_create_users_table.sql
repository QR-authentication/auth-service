-- +goose Up
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE TABLE IF NOT EXISTS users
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    card_number TEXT NOT NULL,
    name        TEXT NOT NULL,
    surname     TEXT NOT NULL,
    password    TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS users;
DROP EXTENSION IF EXISTS pgcrypto;
