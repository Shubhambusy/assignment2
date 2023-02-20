package routes

import (
	"apiass/controller"
	"apiass/helper"
	"apiass/service"
	_ "fmt"

	"github.com/gin-gonic/gin"
)

var (
	bankService    service.BankService       = service.NewBankService()
	BankController controller.BankController = controller.NewBankController(bankService)
)

func NewBank(ctx *gin.Context) {
	bank, err := BankController.Save(ctx)
	helper.HandleResponse(ctx, bank, err)
}

func ViewBank(ctx *gin.Context) {
	bank, err := BankController.Find(ctx)
	helper.HandleResponse(ctx, bank, err)
}

func UpdateBank(ctx *gin.Context) {
	bank, err := BankController.Update(ctx)
	helper.HandleResponse(ctx, bank, err)
}

func DeleteBank(ctx *gin.Context) {
	err := BankController.Delete(ctx)
	helper.HandleResponse(ctx, gin.H{"message": "Bank deleted succesfully"}, err)
}
