package routes

import (
	"apiass/controller"
	"apiass/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	accountService    service.AccountService       = service.NewAccountService()
	AccountController controller.AccountController = controller.NewAccountController(accountService)
)

func NewAccount(ctx *gin.Context) {	
	account, err := AccountController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, account)
	}
}

func NewJointAccount(ctx *gin.Context) {	
	account, err := AccountController.SaveJointAccount(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, account)
	}
}

func ViewAccount(ctx *gin.Context) {
	account, err := AccountController.Find(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, account)
	}
}

func ViewAccountTransections(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "comming soon"})
}

func DeleteAccount(ctx *gin.Context) {
	err := AccountController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Account deleted succesfully"})
	}
}
