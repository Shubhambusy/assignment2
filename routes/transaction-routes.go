package routes

import (
	"apiass/controller"
	"apiass/helper"
	"apiass/service"
	
	"github.com/gin-gonic/gin"
)

var (
	transactionService    service.TransactionService       = service.NewTransactionService()
	TransactionController controller.TransactionController = controller.NewTransactionController(transactionService)
)

func PerformTransaction(ctx *gin.Context) {
	transaction, err := TransactionController.Save(ctx)
	helper.HandleResponse(ctx, transaction, err)
}

func ViewTransaction(ctx *gin.Context) {
	transaction, err := TransactionController.Find(ctx)
	helper.HandleResponse(ctx, transaction, err)
}

func ViewAccountTransactions(ctx *gin.Context) {
	transactions, err := TransactionController.FindByAccount(ctx)
	helper.HandleResponse(ctx, transactions, err)
}
