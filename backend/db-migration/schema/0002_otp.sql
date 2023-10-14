-- +goose Up
CREATE TABLE otp(
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id) NOT NULL,
    secret TEXT NOT NULL DEFAULT '',
    is_verified BOOLEAN NOT NULL DEFAULT FALSE,
    expiry_time TIMESTAMPTZ NOT NULL DEFAULT (NOW()),
    verfied_date TIMESTAMPTZ NOT NULL DEFAULT (NOW()),
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW()),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT (NOW())    
);

-- +goose Down
DROP TABLE otp;