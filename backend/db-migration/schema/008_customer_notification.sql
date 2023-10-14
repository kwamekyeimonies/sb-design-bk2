-- +goose Up
CREATE TABLE customer_notification(
    id UUID PRIMARY KEY,
    customer_id UUID NOT NULL REFERENCES customer(id),
    user_id UUID NOT NULL REFERENCES users(id),
    customer_account_id UUID NOT NULL REFERENCES customer_account(id),
    transaction_type TEXT NOT NULL DEFAULT '',
    message TEXT NOT NULL DEFAULT '',
    amount TEXT NOT NULL DEFAULT '',
    current_amount TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT False,
    deleted_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE customer_notification;