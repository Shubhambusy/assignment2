package routes

import (
	"apiass/controller"
	"apiass/helper"
	"apiass/service"
	_ "fmt"

	"github.com/gin-gonic/gin"
)

var(
	customerService service.CustomerService = service.NewCustomerService()
	CustomerController controller.CustomerController = controller.NewCustomerController(customerService)
)

func NewCustomer (ctx *gin.Context) {
	customer, err := CustomerController.Save(ctx)
	helper.HandleResponse(ctx, customer, err)
}

func ViewCustomer (ctx *gin.Context) {
	customer, err := CustomerController.Find(ctx)
	helper.HandleResponse(ctx, customer, err)
}

func UpdateCustomer (ctx *gin.Context) {
	customer, err := CustomerController.Update(ctx)
	helper.HandleResponse(ctx, customer, err)
}

func DeleteCustomer (ctx *gin.Context) {
	err := CustomerController.Delete(ctx)
	helper.HandleResponse(ctx,  gin.H{"message":"customer deleted succesfully"}, err)
}