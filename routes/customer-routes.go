package routes

import (
	"apiass/controller"
	"apiass/service"
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var(
	customerService service.CustomerService = service.NewCustomerService()
	CustomerController controller.CustomerController = controller.NewCustomerController(customerService)
)

func NewCustomer (ctx *gin.Context) {
	customer, err := CustomerController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, customer)
	}
}

func ViewCustomer (ctx *gin.Context) {
	customer, err := CustomerController.Find(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, customer)
	}
}

func UpdateCustomer (ctx *gin.Context) {
	customer, err := CustomerController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, customer)
	}
}

func DeleteCustomer (ctx *gin.Context) {
	err := CustomerController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message":"customer deleted succesfully"})
	}
}