package user

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/kwamekyeimonies/sb-design-bk/config"
	"github.com/kwamekyeimonies/sb-design-bk/db-migration/db-utils"
	"github.com/kwamekyeimonies/sb-design-bk/models"
	"github.com/kwamekyeimonies/sb-design-bk/utils"
)

func (userS *UserRepositoryImpl) CreateUserAccount(user *models.User, ctx context.Context) (*db.User, error) {

	_, err := userS.pg.PostgresDbApi().DB.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return nil, errors.New("email already exist")
	}
	_, err = userS.pg.PostgresDbApi().DB.GetUserByPhoneNumber(ctx, user.PhoneNumber)
	if err == nil {
		return nil, errors.New("phone-number already exist")
	}

	config, err := config.LoadInitializer("")
	if err != nil {
		return nil, err
	}

	hashedPassword, err := utils.HashPasswordHandler(user.Password, config.PASSWORD_KEY)
	if err != nil {
		return nil, err
	}

	dbParams := &db.CreateUsersParams{
		ID:          uuid.New(),
		FullName:    user.FullName,
		Email:       user.Email,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
		DeletedAt:   time.Now(),
		Dob:         user.Dob,
		PhoneNumber: user.PhoneNumber,
		ProfilePic:  user.ProfilePic,
		UserStatus:  false,
		Password:    hashedPassword,
		Role:        user.Role,
		BankBranch:  user.BankBranch,
		IsVerified:  false,
	}

	dbResponse, err := userS.pg.PostgresDbApi().DB.CreateUsers(ctx, *dbParams)
	if err != nil {
		return nil, err
	}

	return &dbResponse, nil

}
