package models

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
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
	AccountName            string    `json:"account_name"`
	AccountNumber          string    `json:"account_number"`
	AccountType            string    `json:"account_type"`
	Description            string    `json:"description"`
	LoanAmount             string    `json:"loan_amount"`
	Interest               string    `json:"interest"`
	DueDate                string    `json:"due_date"`
	Principal              string    `json:"principal"`
	Loantime               string    `json:"loan_time"`
	Rate                   string    `json:"rate"`
	RemainingAmount        string    `json:"remaining_amount"`
	AmountToBePaid         string    `json:"amount_to_be_paid"`
	LoanBalance            string    `json:"loan_balance"`
	InitialLoan            string    `json:"initial_loan"`
	CurrentLoan            string    `json:"current_amount"`
	PaidAmount             string    `json:"paid_amount"`
	UserId                 uuid.UUID `json:"user_id"`
}
