-- name: CreateUsers :one
INSERT INTO users(
    id,
    full_name,
    email,
    created_at,
    updated_at,
    is_deleted,
    deleted_at,
    dob,
    phone_number,
    profile_pic,
    user_status,
    password,
    role,
    bank_branch,
    is_verified,
    api_key
    
)VALUES(
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,
    encode(sha256(random()::text::bytea),'hex')
)RETURNING *;

-- name: GetAllUsers :many
SELECT id,full_name, email,created_at,updated_at,phone_number,profile_pic,role,api_key 
FROM users;

-- name: GetUserById :one
SELECT id,full_name, email,created_at,updated_at,phone_number,profile_pic,role,api_key ,password,bank_branch,dob
FROM users WHERE id = $1 AND is_deleted=False;

-- name: GetUserByPhoneNumber :one
SELECT id,full_name, email,created_at,updated_at,phone_number,profile_pic,role,api_key,password ,bank_branch,dob
FROM users WHERE phone_number = $1 AND is_deleted=False;

-- name: GetUserByEmail :one
SELECT id,full_name, email,created_at,updated_at,phone_number,profile_pic,role,api_key,password,is_verified,bank_branch ,dob
FROM users WHERE email = $1 AND is_deleted=False;

-- name: DeleteUser :exec
UPDATE users
SET
   is_deleted = True,
   deleted_at = NOW()
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET
    updated_at = NOW(),
    phone_number = $1,
    profile_pic = $2,
    user_status = $3,
    password = $4,
    api_key = encode(sha256(random()::text::bytea),'hex')
WHERE id = $5;

-- name: UpdateUserbyAdmin :exec
UPDATE users
SET
    updated_at = NOW(),
    phone_number = $1,
    profile_pic = $2,
    user_status = $3,
    password = $4,
    role = $5,
    api_key = encode(sha256(random()::text::bytea),'hex')
WHERE id = $6;

-- name: ForgetPassword :exec
UPDATE users
SET
    updated_at = NOW(),
    password = $1,
    api_key = encode(sha256(random()::text::bytea),'hex')
WHERE id = $2;

-- name: UserStatusUpdate :exec
UPDATE users
SET
    updated_at = NOW(),
    is_verified = $1,
    api_key = encode(sha256(random()::text::bytea),'hex')
WHERE id = $2;

-- name: UpdateUserPhoneNumber :exec
UPDATE users
SET
    updated_at = NOW(),
    phone_number = $1,
    api_key = encode(sha256(random()::text::bytea),'hex')
WHERE id = $2;