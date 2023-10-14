-- +goose Up
CREATE TABLE customer_account(
    id UUID PRIMARY KEY,
    account_name TEXT NOT NULL,
    account_number TEXT NOT NULL,
    account_type TEXT NOT NULL,
    currency_type TEXT NOT NULL,
    loan_balance TEXT NOT NULL DEFAULT '',
    initial_loan TEXT NOT NULL DEFAULT '',
    current_loan TEXT NOT NULL DEFAULT '',
    amount_to_be_paid TEXT NOT NULL DEFAULT '',
    customer_id UUID REFERENCES customer(id) NOT NULL,
     user_id UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT False,
    deleted_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    phone_number TEXT NOT NULL DEFAULT ''
);

-- +goose Down
DROP TABLE customer_account;