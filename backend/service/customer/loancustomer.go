package customer

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/kwamekyeimonies/sb-design-bk/db-migration/db-utils"
	"github.com/kwamekyeimonies/sb-design-bk/models"
)

func (customerS *CustomerServiceImpl) LoanMoneyToCustomer(customerAcc *models.CustomerLoans, ctx context.Context) (*string, error) {
	_, err := customerS.pg.PostgresDbApi().DB.GetCustomerAccount(ctx, customerAcc.CustomerID)
	if err != nil {
		return nil, errors.New("Customer Account does not exist")
	}

	_, err = customerS.pg.PostgresDbApi().DB.CreateCustomerLoan(ctx, db.CreateCustomerLoanParams{
		ID:              uuid.New(),
		CustomerID:      customerAcc.CustomerID,
		LoanAmount:      customerAcc.LoanAmount,
		PaidAmount:      customerAcc.PaidAmount,
		RemainingAmount: customerAcc.RemainingAmount,
		Principal:       customerAcc.Principal,
		LoanTime:        customerAcc.LoanTime,
		Rate:            customerAcc.Rate,
		Interest:        customerAcc.Interest,
		DueDate:         customerAcc.DueDate,
		UserID:          customerAcc.UserID,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		IsDeleted:       false,
		DeletedAt:       time.Now(),
	})
	if err != nil {
		return nil, err
	}

	msg := "Loan creation successfull"

	return &msg, nil
}
