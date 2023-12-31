package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kwamekyeimonies/sb-design-bk/models"
)

func (customerControl *CustomerController) CreateCustomerController(ctx *gin.Context) {
	if ctx.Request.ContentLength == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid/empty body request",
		})
		return
	}

	var customer *models.Customer

	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	dbResponse, msg, err := customerControl.customr.CreateAccount(customer, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  msg,
		"response": dbResponse,
	})

}
