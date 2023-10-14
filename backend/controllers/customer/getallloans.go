package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (customerControl *CustomerController) GetAllCustomerLoansController(ctx *gin.Context) {

	dbResponse, err := customerControl.customr.GetCustomersLoanDetails(ctx)
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
