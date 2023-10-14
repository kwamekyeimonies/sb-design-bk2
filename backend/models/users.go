package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsDeleted   bool      `json:"is_deleted"`
	DeletedAt   time.Time `json:"deleted_at"`
	Dob         string    `json:"dob"`
	PhoneNumber string    `json:"phone_number"`
	ProfilePic  string    `json:"profile_pic"`
	UserStatus  bool      `json:"user_status"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	BankBranch  string    `json:"bank_branch"`
	IsVerified  bool      `json:"is_verified"`
}
