-- +goose Up
SELECT * FROM pg_extension WHERE extname = 'uuid-ossp';
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users(
    id UUID PRIMARY KEY,
    full_name TEXT NOT NULL DEFAULT '',
    email TEXT NOT NULL UNIQUE,
    role TEXT NOT NULL DEFAULT 'User',
    password TEXT NOT NULL,
    bank_branch TEXT NOT NULL DEFAULT '',
    dob TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT False,
    deleted_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    phone_number TEXT NOT NULL UNIQUE DEFAULT '',
    profile_pic TEXT NOT NULL DEFAULT '',
    user_status BOOLEAN NOT NULL DEFAULT True,
    is_verified BOOLEAN NOT NULL DEFAULT False,
    api_key TEXT NOT NULL
);

-- +goose Down
DROP TABLE admin;