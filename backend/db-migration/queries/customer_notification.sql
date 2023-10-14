-- name: CreateCustomerNotification :one
INSERT INTO  customer_notification(
    id,
    customer_id,
    user_id,
    customer_account_id,
    transaction_type,
    message,
    amount,
    current_amount,
    created_at,
    updated_at,
    is_deleted,
    deleted_at
)VALUES(
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12
)RETURNING *;

