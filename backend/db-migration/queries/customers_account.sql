-- name: CreateCustomerAccount :one
INSERT INTO customer_account(
    id,
    account_name,
    account_number,
    account_type,
    loan_balance,
    initial_loan,
    current_loan,
    customer_id,
    created_at,
    updated_at,
    is_deleted,
    deleted_at,
    currency_type,
    amount_to_be_paid,
    user_id,
    phone_number
)VALUES(
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16
)RETURNING *;


-- name: GetCustomerAccount :one
SELECT * FROM customer_account WHERE id = $1;

-- name: UpdateCustomerAccount :exec
UPDATE customer_account SET
    loan_balance= $1,
    initial_loan= $2,
    current_loan= $3,
    updated_at = NOW()
WHERE id = $4;
