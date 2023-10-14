package router

import (
	"github.com/gin-gonic/gin"

	"github.com/kwamekyeimonies/sb-design-bk/controllers/customer"
	"github.com/kwamekyeimonies/sb-design-bk/controllers/users"
)

func NonAuthenticatedRoutes(routes *gin.RouterGroup) {
	userRoutes := users.UserController{}

	routes.POST("/api/v1/users/signup", userRoutes.CreateNewUserController)
	routes.POST("/api/v1/users/login", userRoutes.LoginUserController)

}

func AuthenticatedRoutes(routes *gin.RouterGroup) {
	// userRoutes := users.UserController{}
	customerRoutes := customer.CustomerController{}

	routes.POST("/customer/create", customerRoutes.CreateCustomerController)
	routes.POST("/customer/loan", customerRoutes.LoanCustomerController)
	routes.POST("/customer/account", customerRoutes.GetCustomerAccountController)
	routes.POST("/customer/loan/create", customerRoutes.CreateCustomerLoanController)
	routes.GET("/customer/loans", customerRoutes.GetAllCustomerLoansController)

}
