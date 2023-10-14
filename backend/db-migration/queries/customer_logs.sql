-- name: CreateCustomerLogs :one
INSERT INTO customer_logs(
    id,
    customer_id,
    user_id,
    customer_account_id,
    transaction_type,
    created_at,
    updated_at,
    is_deleted,
    deleted_at
)VALUES(
    $1,$2,$3,$4,$5,$6,$7,$8,$9
)RETURNING *;

