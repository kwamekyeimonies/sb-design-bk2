-- +goose Up
CREATE TABLE customer_loans(
    id UUID PRIMARY KEY,
    customer_id UUID NOT NULL REFERENCES customer(id),
    loan_amount TEXT NOT NULL DEFAULT '',
    paid_amount TEXT NOT NULL DEFAULT '',
    interest TEXT NOT NULL DEFAULT '',
    principal TEXT NOT NULL DEFAULT '',
    loan_time TEXT NOT NULL DEFAULT '',
    amount_to_be_paid TEXT NOT NULL DEFAULT '',
    rate TEXT NOT NULL DEFAULT '',
    due_date TEXT NOT NULL DEFAULT '',
    remaining_amount TEXT NOT NULL DEFAULT '',
    user_id UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT False,
    deleted_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE customer_loans;