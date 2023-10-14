package users

import "github.com/kwamekyeimonies/sb-design-bk/service/user"

type UserController struct {
	userCntrl user.UserRepositoryImpl
}

func NewUserController(userCntrl user.UserRepositoryImpl) UserController {
	return UserController{
		userCntrl: userCntrl,
	}
}
