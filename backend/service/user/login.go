package user

import (
	"context"
	"errors"

	"github.com/kwamekyeimonies/sb-design-bk/config"
	"github.com/kwamekyeimonies/sb-design-bk/db-migration/db-utils"
	"github.com/kwamekyeimonies/sb-design-bk/models"
	"github.com/kwamekyeimonies/sb-design-bk/utils"
)

func (userS *UserRepositoryImpl) LoginUserAccount(user *models.User, ctx context.Context) (*db.GetUserByEmailRow, error) {
	data, err := userS.pg.PostgresDbApi().DB.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, errors.New("email does not exist")
	}

	config, err := config.LoadInitializer("")
	if err != nil {
		return nil, err
	}

	passwordChecker := utils.CheckPasswordHashHandler(user.Password, data.Password, config.PASSWORD_KEY)
	if passwordChecker != nil {
		return nil, errors.New("invalid credentials")
	}

	return &data, nil

}
