package customer

import (
	"context"

	"github.com/kwamekyeimonies/sb-design-bk/db-migration/db-utils"
	"github.com/kwamekyeimonies/sb-design-bk/models"
)

func (customerS *CustomerServiceImpl) GetCustomerAccountDetails(customer *models.Customer, ctx context.Context) (*db.CustomerAccount, error) {
	dbResponse, err := customerS.pg.PostgresDbApi().DB.GetCustomerAccount(ctx, customer.ID)
	if err != nil {
		return nil, err
	}

	return &dbResponse, nil
}
