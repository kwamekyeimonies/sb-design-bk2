// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

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
}

type CustomerAccount struct {
	ID             uuid.UUID `json:"id"`
	AccountName    string    `json:"account_name"`
	AccountNumber  string    `json:"account_number"`
	AccountType    string    `json:"account_type"`
	CurrencyType   string    `json:"currency_type"`
	LoanBalance    string    `json:"loan_balance"`
	InitialLoan    string    `json:"initial_loan"`
	CurrentLoan    string    `json:"current_loan"`
	AmountToBePaid string    `json:"amount_to_be_paid"`
	CustomerID     uuid.UUID `json:"customer_id"`
	UserID         uuid.UUID `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	IsDeleted      bool      `json:"is_deleted"`
	DeletedAt      time.Time `json:"deleted_at"`
	PhoneNumber    string    `json:"phone_number"`
}

type CustomerLoan struct {
	ID              uuid.UUID `json:"id"`
	CustomerID      uuid.UUID `json:"customer_id"`
	LoanAmount      string    `json:"loan_amount"`
	PaidAmount      string    `json:"paid_amount"`
	Interest        string    `json:"interest"`
	Principal       string    `json:"principal"`
	LoanTime        string    `json:"loan_time"`
	AmountToBePaid  string    `json:"amount_to_be_paid"`
	Rate            string    `json:"rate"`
	DueDate         string    `json:"due_date"`
	RemainingAmount string    `json:"remaining_amount"`
	UserID          uuid.UUID `json:"user_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	IsDeleted       bool      `json:"is_deleted"`
	DeletedAt       time.Time `json:"deleted_at"`
}

type CustomerLog struct {
	ID                uuid.UUID `json:"id"`
	CustomerID        uuid.UUID `json:"customer_id"`
	UserID            uuid.UUID `json:"user_id"`
	CustomerAccountID uuid.UUID `json:"customer_account_id"`
	TransactionType   string    `json:"transaction_type"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	IsDeleted         bool      `json:"is_deleted"`
	DeletedAt         time.Time `json:"deleted_at"`
}

type CustomerNotification struct {
	ID                uuid.UUID `json:"id"`
	CustomerID        uuid.UUID `json:"customer_id"`
	UserID            uuid.UUID `json:"user_id"`
	CustomerAccountID uuid.UUID `json:"customer_account_id"`
	TransactionType   string    `json:"transaction_type"`
	Message           string    `json:"message"`
	Amount            string    `json:"amount"`
	CurrentAmount     string    `json:"current_amount"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	IsDeleted         bool      `json:"is_deleted"`
	DeletedAt         time.Time `json:"deleted_at"`
}

type Otp struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Secret      string    `json:"secret"`
	IsVerified  bool      `json:"is_verified"`
	ExpiryTime  time.Time `json:"expiry_time"`
	VerfiedDate time.Time `json:"verfied_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Transaction struct {
	ID                uuid.UUID `json:"id"`
	CustomerID        uuid.UUID `json:"customer_id"`
	UserID            uuid.UUID `json:"user_id"`
	CustomerAccountID uuid.UUID `json:"customer_account_id"`
	Description       string    `json:"description"`
	InitialAmount     string    `json:"initial_amount"`
	FinalAmount       string    `json:"final_amount"`
	TransactionType   string    `json:"transaction_type"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	IsDeleted         bool      `json:"is_deleted"`
	DeletedAt         time.Time `json:"deleted_at"`
}

type User struct {
	ID          uuid.UUID `json:"id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	Password    string    `json:"password"`
	BankBranch  string    `json:"bank_branch"`
	Dob         string    `json:"dob"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsDeleted   bool      `json:"is_deleted"`
	DeletedAt   time.Time `json:"deleted_at"`
	PhoneNumber string    `json:"phone_number"`
	ProfilePic  string    `json:"profile_pic"`
	UserStatus  bool      `json:"user_status"`
	IsVerified  bool      `json:"is_verified"`
	ApiKey      string    `json:"api_key"`
}

type UsersLog struct {
	ID                 uuid.UUID `json:"id"`
	CustomerID         uuid.UUID `json:"customer_id"`
	UserID             uuid.UUID `json:"user_id"`
	CustomerAccountID  uuid.UUID `json:"customer_account_id"`
	TransactionType    string    `json:"transaction_type"`
	TransactionPurpose string    `json:"transaction_purpose"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	IsDeleted          bool      `json:"is_deleted"`
	DeletedAt          time.Time `json:"deleted_at"`
}
