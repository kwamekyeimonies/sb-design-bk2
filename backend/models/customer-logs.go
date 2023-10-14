package models

import (
	"time"

	"github.com/google/uuid"
)

type CustomerLogs struct {
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
