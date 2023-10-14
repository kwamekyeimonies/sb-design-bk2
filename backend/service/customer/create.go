package customer

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/kwamekyeimonies/sb-design-bk/db-migration/db-utils"
	"github.com/kwamekyeimonies/sb-design-bk/models"
)

func (customerS *CustomerServiceImpl) CreateAccount(customer *models.Customer, ctx context.Context) (*db.Customer, *string, error) {
	dbResponse, err := customerS.pg.PostgresDbApi().DB.CreateCustomer(ctx, db.CreateCustomerParams{
		ID:                     uuid.New(),
		FullName:               customer.FullName,
		Email:                  customer.Email,
		Role:                   "customer",
		Password:               "demo",
		Organization:           "Bank",
		MaritalStatus:          customer.MaritalStatus,
		EmploymentStatus:       customer.EmploymentStatus,
		NameOfEmployer:         customer.NameOfEmployer,
		DateOfBirth:            customer.DateOfBirth,
		IdentificationCard:     customer.IdentificationCard,
		IdentificationCardType: customer.IdentificationCardType,
		Address:                customer.Address,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
		IsDeleted:              false,
		DeletedAt:              time.Now(),
		PhoneNumber:            customer.PhoneNumber,
		ProfilePic:             customer.ProfilePic,
		UserStatus:             false,
		IsVerified:             false,
	})

	if err != nil {
		return nil, nil, err
	}

	// accountResponse, err := customerS.pg.PostgresDbApi().DB.CreateCustomerAccount(ctx, db.CreateCustomerAccountParams{
	// 	ID:            uuid.New(),
	// 	AccountName:   dbResponse.FullName,
	// 	AccountNumber: dbResponse.ID.String(),
	// 	AccountType:   "Loan Account",
	// 	CurrencyType:  "GHC.",
	// 	LoanBalance:   customer.LoanBalance,
	// 	InitialLoan:   customer.InitialLoan,
	// 	CurrentLoan:   customer.CurrentLoan,
	// 	CustomerID:    dbResponse.ID,
	// 	UserID:        customer.UserId,
	// 	CreatedAt:     time.Now(),
	// 	UpdatedAt:     time.Now(),
	// 	IsDeleted:     false,
	// 	DeletedAt:     time.Now(),
	// 	PhoneNumber:   customer.PhoneNumber,
	// })

	// if err != nil {
	// 	return nil, nil, nil, err
	// }

	msg := "Customer account created successfully"

	return &dbResponse, &msg, nil

}
