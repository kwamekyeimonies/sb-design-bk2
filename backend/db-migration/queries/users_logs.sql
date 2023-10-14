-- name: CreateUserLogs :one
INSERT INTO users_logs(
    id,
    customer_id,
    user_id,
    customer_account_id,
    transaction_type,
    transaction_purpose,
    created_at,
    updated_at,
    is_deleted,
    deleted_at
)VALUES(
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10
)RETURNING *;