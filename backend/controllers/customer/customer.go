package customer

import "github.com/kwamekyeimonies/sb-design-bk/service/customer"

type CustomerController struct {
	customr customer.CustomerServiceImpl
}

func NewCustomerController(customr customer.CustomerServiceImpl) CustomerController {
	return CustomerController{
		customr: customr,
	}
}
