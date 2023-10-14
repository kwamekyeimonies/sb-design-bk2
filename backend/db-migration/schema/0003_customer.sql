-- +goose Up
SELECT * FROM pg_extension WHERE extname = 'uuid-ossp';
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE customer
(
    id UUID PRIMARY KEY,
    full_name TEXT NOT NULL DEFAULT '',
    email TEXT NOT NULL UNIQUE,
    role TEXT NOT NULL DEFAULT 'User',
    password TEXT NOT NULL,
    organization TEXT NOT NULL DEFAULT '',
    marital_status TEXT NOT NULL DEFAULT 'Single',
    employment_status TEXT NOT NULL DEFAULT '',
    name_of_employer TEXT NOT NULL DEFAULT '',
    date_of_birth TEXT NOT NULL DEFAULT '',
    identification_card TEXT NOT NULL DEFAULT '',
    identification_card_type TEXT NOT NULL DEFAULT '',
    address TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT False,
    deleted_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    phone_number TEXT NOT NULL UNIQUE DEFAULT '',
    profile_pic TEXT NOT NULL DEFAULT '',
    user_status BOOLEAN NOT NULL DEFAULT True,
    is_verified BOOLEAN NOT NULL DEFAULT False
  
);

-- +goose Down
DROP TABLE admin;