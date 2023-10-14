-- name: CreateCustomerLoan :one
INSERT INTO customer_loans(
    id,
    customer_id,
    loan_amount,
    paid_amount,
    remaining_amount,
    amount_to_be_paid,
    principal,
    loan_time,
    rate,
    interest,
    due_date,
    user_id,
    created_at,
    updated_at,
    is_deleted,
    deleted_at
)VALUES(
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16
)RETURNING *;



-- name: UpdateLoanTransaction :exec
UPDATE customer_loans SET
    paid_amount = $1,
    loan_amount = $2,
    remaining_amount = $3,
    user_id = $4,
    principal = $5,
    interest = $6,
    due_date = $7,
    rate = $8,
    updated_at = NOW()
WHERE customer_id = $9;
    


-- name: GetLoanByCustomerId :many
SELECT * FROM customer_loans WHERE customer_id = $1 AND is_deleted = False;

-- name: GetAllLoans :many
SELECT * FROM customer_loans;

