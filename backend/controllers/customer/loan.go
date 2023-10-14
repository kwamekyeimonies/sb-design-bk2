package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kwamekyeimonies/sb-design-bk/models"
)

func (customerControl *CustomerController) LoanCustomerController(ctx *gin.Context) {
	if ctx.Request.ContentLength == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid/empty body request",
		})
		return
	}

	var customer *models.CustomerLoans

	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	dbResponse, err := customerControl.customr.LoanMoneyToCustomer(customer, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": dbResponse,
	})

}
