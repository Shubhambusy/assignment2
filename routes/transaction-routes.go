package routes

import (
	"apiass/controller"
	"apiass/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	transactionService    service.TransactionService       = service.NewTransactionService()
	TransactionController controller.TransactionController = controller.NewTransactionController(transactionService)
)

func PerformTransaction(ctx *gin.Context) {
	transaction, err := TransactionController.Save(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, transaction)
	}
}

func ViewTransaction(ctx *gin.Context) {
	transaction, err := TransactionController.Find(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, transaction)
	}
}
