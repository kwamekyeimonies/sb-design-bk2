-- name: CreateOTP :one
INSERT INTO otp(
    id,
    user_id,
   is_verified,
    verfied_date,
    created_at,
    updated_at,
    expiry_time,
    secret
)VALUES(
    $1,$2,$3,$4,$5,$6,$7,$8
)RETURNING *;

-- name: GetOTPByUserID :one
SELECT * FROM otp WHERE user_id = $1;

-- name: UpdateOTP :exec
UPDATE otp
SET 
    secret = $1,
    expiry_time=$2 ,
     is_verified = $3
WHERE id = $4;

