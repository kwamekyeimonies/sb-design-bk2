package customer

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/kwamekyeimonies/sb-design-bk/db-migration/db-utils"
	"github.com/kwamekyeimonies/sb-design-bk/models"
	"github.com/kwamekyeimonies/sb-design-bk/utils"
)

func (customerS *CustomerServiceImpl) CreateLoanAccount(customer *models.Customer, ctx context.Context) (*db.CustomerAccount, *db.CustomerLoan, *string, error) {

	dbResponse, err := customerS.pg.PostgresDbApi().DB.GetCustomerAccountById(ctx, customer.ID)
	if err != nil {
		return nil, nil, nil, errors.New("Customer not eligible, needs to have an accout")
	}

	principal, err := utils.ConvertToDouble(customer.Principal)
	if err != nil {
		return nil, nil, nil, err
	}

	rate, err := utils.ConvertToDouble(customer.Rate)
	if err != nil {
		return nil, nil, nil, err
	}

	loantime, err := utils.ConvertToDouble(customer.Loantime)
	if err != nil {
		return nil, nil, nil, err
	}

	amountToBepaid := utils.SimpleInterestCalculator(principal, rate, loantime)

	totalAmount := amountToBepaid + principal
	totalAmountString := utils.ConvertToString(totalAmount)
	loanAmunt, _ := utils.ConvertToDouble(customer.CurrentLoan)

	currentloan := loanAmunt - amountToBepaid
	currentLoanString := utils.ConvertToString(currentloan)

	accountResponse, err := customerS.pg.PostgresDbApi().DB.CreateCustomerAccount(ctx, db.CreateCustomerAccountParams{
		ID:             uuid.New(),
		AccountName:    dbResponse.FullName,
		AccountNumber:  dbResponse.ID.String(),
		AccountType:    "Loan Account",
		LoanBalance:    currentLoanString,
		InitialLoan:    utils.ConvertToString(loanAmunt),
		CurrentLoan:    currentLoanString,
		CustomerID:     dbResponse.ID,
		CreatedAt:      time.Now(),
		IsDeleted:      false,
		DeletedAt:      time.Now(),
		CurrencyType:   "GHC.",
		AmountToBePaid: totalAmountString,
		UserID:         customer.UserId,
		PhoneNumber:    dbResponse.PhoneNumber,
	})

	if err != nil {
		return nil, nil, nil, err
	}

	loanResponse, err := customerS.pg.PostgresDbApi().DB.CreateCustomerLoan(ctx, db.CreateCustomerLoanParams{
		ID:              uuid.New(),
		CustomerID:      dbResponse.ID,
		LoanAmount:      customer.Principal,
		PaidAmount:      customer.PaidAmount,
		RemainingAmount: currentLoanString,
		AmountToBePaid:  totalAmountString,
		Principal:       customer.Principal,
		LoanTime:        customer.Loantime,
		Rate:            customer.Rate,
		Interest:        customer.Interest,
		DueDate:         customer.DueDate,
		UserID:          customer.UserId,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		IsDeleted:       false,
		DeletedAt:       time.Now(),
	})
	if err != nil {
		return nil, nil, nil, err
	}

	msg := "Customer account created successfully"

	return &accountResponse, &loanResponse, &msg, nil

}
