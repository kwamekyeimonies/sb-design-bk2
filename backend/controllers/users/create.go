package users

import (
	"github.com/gin-gonic/gin"

	"github.com/kwamekyeimonies/sb-design-bk/models"
)

func (userControl *UserController) CreateNewUserController(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := userControl.userCntrl.CreateUserAccount(&user, ctx)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message":  "user created",
		"response": response,
	})
}
