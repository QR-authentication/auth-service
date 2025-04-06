-- +goose Up
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE TABLE IF NOT EXISTS users
(
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name         TEXT NOT NULL,
    surname      TEXT NOT NULL,
    patronymic   TEXT,
    department   TEXT NOT NULL,
    group_number TEXT,
    card_number  TEXT NOT NULL,
    password     TEXT NOT NULL,
    created_at   TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS users;
DROP EXTENSION IF EXISTS pgcrypto;
