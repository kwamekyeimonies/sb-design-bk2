package customer

import (
	"context"

	"github.com/kwamekyeimonies/sb-design-bk/db-migration/db-utils"
)

func (customerS *CustomerServiceImpl) GetCustomersLoanDetails(ctx context.Context) ([]db.CustomerLoan, error) {
	dbResponse, err := customerS.pg.PostgresDbApi().DB.GetAllLoans(ctx)
	if err != nil {
		return nil, err
	}

	return dbResponse, nil
}
