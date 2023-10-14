package user

import (
	"context"

	"github.com/kwamekyeimonies/sb-design-bk/db-migration/db-utils"
	"github.com/kwamekyeimonies/sb-design-bk/dbapi"
	"github.com/kwamekyeimonies/sb-design-bk/models"
)

type UserRepository interface {
	CreateUserAccount(usr *models.User, ctx context.Context) (*db.User, error)
	// UpdateUserAccount(usr *models.User, ctx context.Context) (*string, error)
	// DeleteUserAccount(usr *models.User, ctx context.Context) (*string, error)
	LoginUserAccount(usr *models.User, ctx context.Context) (*db.GetUserByEmailRow, error)
}

type UserRepositoryImpl struct {
	pg  *dbapi.NewPgApiConfig
	ctx context.Context
}

func NewUserRepository(pg *dbapi.NewPgApiConfig, ctx context.Context) UserRepository {
	return &UserRepositoryImpl{
		pg:  pg,
		ctx: ctx,
	}
}
