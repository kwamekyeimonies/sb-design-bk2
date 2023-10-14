package customer

import (
	"context"

	"github.com/kwamekyeimonies/sb-design-bk/db-migration/db-utils"
	"github.com/kwamekyeimonies/sb-design-bk/dbapi"
	"github.com/kwamekyeimonies/sb-design-bk/models"
)

type CustomerRespository interface {
	CreateAccount(customer *models.Customer, ctx context.Context) (*db.Customer, *string, error)
	LoanMoneyToCustomer(customerAcc *models.CustomerLoans, ctx context.Context) (*string, error)
	GetCustomerAccountDetails(customer *models.Customer, ctx context.Context) (*db.CustomerAccount, error)
	GetCustomersLoanDetails(ctx context.Context) ([]db.CustomerLoan, error)
}

type CustomerServiceImpl struct {
	pg  *dbapi.NewPgApiConfig
	ctx context.Context
}

func NewCustomerRespositoryImpl(pg *dbapi.NewPgApiConfig, ctx context.Context) CustomerRespository {
	return &CustomerServiceImpl{
		pg:  pg,
		ctx: ctx,
	}
}
