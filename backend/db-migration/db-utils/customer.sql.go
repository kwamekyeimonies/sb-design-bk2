// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: customer.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createCustomer = `-- name: CreateCustomer :one
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
)RETURNING id, full_name, email, role, password, organization, marital_status, employment_status, name_of_employer, date_of_birth, identification_card, identification_card_type, address, created_at, updated_at, is_deleted, deleted_at, phone_number, profile_pic, user_status, is_verified
`

type CreateCustomerParams struct {
	ID                     uuid.UUID `json:"id"`
	FullName               string    `json:"full_name"`
	Email                  string    `json:"email"`
	Role                   string    `json:"role"`
	Password               string    `json:"password"`
	Organization           string    `json:"organization"`
	MaritalStatus          string    `json:"marital_status"`
	EmploymentStatus       string    `json:"employment_status"`
	NameOfEmployer         string    `json:"name_of_employer"`
	DateOfBirth            string    `json:"date_of_birth"`
	IdentificationCard     string    `json:"identification_card"`
	IdentificationCardType string    `json:"identification_card_type"`
	Address                string    `json:"address"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
	IsDeleted              bool      `json:"is_deleted"`
	DeletedAt              time.Time `json:"deleted_at"`
	PhoneNumber            string    `json:"phone_number"`
	ProfilePic             string    `json:"profile_pic"`
	UserStatus             bool      `json:"user_status"`
	IsVerified             bool      `json:"is_verified"`
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error) {
	row := q.db.QueryRow(ctx, createCustomer,
		arg.ID,
		arg.FullName,
		arg.Email,
		arg.Role,
		arg.Password,
		arg.Organization,
		arg.MaritalStatus,
		arg.EmploymentStatus,
		arg.NameOfEmployer,
		arg.DateOfBirth,
		arg.IdentificationCard,
		arg.IdentificationCardType,
		arg.Address,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.IsDeleted,
		arg.DeletedAt,
		arg.PhoneNumber,
		arg.ProfilePic,
		arg.UserStatus,
		arg.IsVerified,
	)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Role,
		&i.Password,
		&i.Organization,
		&i.MaritalStatus,
		&i.EmploymentStatus,
		&i.NameOfEmployer,
		&i.DateOfBirth,
		&i.IdentificationCard,
		&i.IdentificationCardType,
		&i.Address,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.DeletedAt,
		&i.PhoneNumber,
		&i.ProfilePic,
		&i.UserStatus,
		&i.IsVerified,
	)
	return i, err
}

const getCustomerAccountById = `-- name: GetCustomerAccountById :one
SELECT id, full_name, email, role, password, organization, marital_status, employment_status, name_of_employer, date_of_birth, identification_card, identification_card_type, address, created_at, updated_at, is_deleted, deleted_at, phone_number, profile_pic, user_status, is_verified FROM customer WHERE id = $1
`

func (q *Queries) GetCustomerAccountById(ctx context.Context, id uuid.UUID) (Customer, error) {
	row := q.db.QueryRow(ctx, getCustomerAccountById, id)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Role,
		&i.Password,
		&i.Organization,
		&i.MaritalStatus,
		&i.EmploymentStatus,
		&i.NameOfEmployer,
		&i.DateOfBirth,
		&i.IdentificationCard,
		&i.IdentificationCardType,
		&i.Address,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.DeletedAt,
		&i.PhoneNumber,
		&i.ProfilePic,
		&i.UserStatus,
		&i.IsVerified,
	)
	return i, err
}