package routes

import (
	"apiass/controller"
	"apiass/service"
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var(
	bankService service.BankService = service.NewBankService()
	BankController controller.BankController = controller.NewBankController(bankService)
)

func NewBank (ctx *gin.Context) {
	bank, err := BankController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, bank)
	}
}

func ViewBank (ctx *gin.Context) {
	bank, err := BankController.Find(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, bank)
	}
}

func UpdateBank (ctx *gin.Context) {
	bank, err := BankController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, bank)
	}
}

func DeleteBank (ctx *gin.Context) {
	err := BankController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message":"Bank deleted succesfully"})
	}
}