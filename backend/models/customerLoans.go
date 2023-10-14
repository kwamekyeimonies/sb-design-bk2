package models

import (
	"time"

	"github.com/google/uuid"
)

type CustomerLoans struct {
	ID              uuid.UUID `json:"id"`
	CustomerID      uuid.UUID `json:"customer_id"`
	LoanAmount      string    `json:"loan_amount"`
	PaidAmount      string    `json:"paid_amount"`
	RemainingAmount string    `json:"remaining_amount"`
	Principal       string    `json:"principal"`
	LoanTime        string    `json:"loan_time"`
	Rate            string    `json:"rate"`
	Interest        string    `json:"interest"`
	DueDate         string    `json:"due_date"`
	UserID          uuid.UUID `json:"user_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	IsDeleted       bool      `json:"is_deleted"`
	DeletedAt       time.Time `json:"deleted_at"`
}
