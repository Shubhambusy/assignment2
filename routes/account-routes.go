package routes

import (
	"apiass/controller"
	"apiass/helper"
	"apiass/service"

	"github.com/gin-gonic/gin"
)

var (
	accountService    service.AccountService       = service.NewAccountService()
	AccountController controller.AccountController = controller.NewAccountController(accountService)
)

func NewAccount(ctx *gin.Context) {
	account, err := AccountController.Save(ctx)
	helper.HandleResponse(ctx, account, err)
}

func NewJointAccount(ctx *gin.Context) {
	account, err := AccountController.SaveJointAccount(ctx)
	helper.HandleResponse(ctx, account, err)
}

func ViewAccount(ctx *gin.Context) {
	account, err := AccountController.Find(ctx)
	helper.HandleResponse(ctx, account, err)
}

func DeleteAccount(ctx *gin.Context) {
	err := AccountController.Delete(ctx)
	helper.HandleResponse(ctx, gin.H{"message": "Account deleted succesfully"}, err)
}
