package models

import (
	"time"

	"github.com/google/uuid"
)

type CustomerAccount struct {
	ID             uuid.UUID `json:"id"`
	AccountName    string    `json:"account_name"`
	AccountNumber  string    `json:"account_number"`
	AccountType    string    `json:"account_type"`
	Description    string    `json:"description"`
	DepsitAmount   string    `json:"deposit_amount"`
	AccountBalance string    `json:"account_balance"`
	InitialDeposit string    `json:"initial_deposit"`
	CurrentAmount  string    `json:"current_amount"`
	CustomerID     uuid.UUID `json:"customer_id"`
	UserId         uuid.UUID `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	IsDeleted      bool      `json:"is_deleted"`
	DeletedAt      time.Time `json:"deleted_at"`
	PhoneNumber    string    `json:"phone_number"`
}
