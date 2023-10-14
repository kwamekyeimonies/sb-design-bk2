-- name: CreateCustomer :one
INSERT INTO customer
(   id,
    full_name,
    email,
    role,
    password,
    organization,
    marital_status,
    employment_status,
    name_of_employer,
    date_of_birth,
    identification_card,
    identification_card_type,
    address,
    created_at,
    updated_at,
    is_deleted,
    deleted_at,
    phone_number,
    profile_pic,
    user_status,
    is_verified
)VALUES(
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21
)RETURNING *;

-- name: GetCustomerAccountById :one
SELECT * FROM customer WHERE id = $1;