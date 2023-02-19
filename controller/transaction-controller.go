package controller

import (
	"apiass/entity"
	"apiass/service"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionController interface {
	Find(ctx *gin.Context) (entity.Transaction, error)
	Save(ctx *gin.Context) (entity.Transaction, error)
}

type transactionController struct {
	service service.TransactionService
}

func NewTransactionController(service service.TransactionService) TransactionController {
	return &transactionController{
		service: service,
	}
}

func (c *transactionController) Find(ctx *gin.Context) (entity.Transaction, error) {
	var transaction entity.Transaction
	Id, err := uuid.Parse(ctx.Params.ByName("id"))
	if err != nil {
		return transaction, err
	}

	transaction, err = c.service.Find(Id)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (c *transactionController) Save(ctx *gin.Context) (entity.Transaction, error) {

	var transaction entity.Transaction

	rawData, err := ctx.GetRawData()
	if (err != nil) {
		return transaction, err
	}
	var body map[string]interface{}

	if err := json.Unmarshal(rawData, &body); err != nil {
		return transaction, err
	}

	transactionType, err := getTransactionType(body)
	if (err != nil) {
		return transaction, err
	}
	if (transactionType == "cash_deposit" || transactionType == "cash_withdraw") {
		transaction, err = c.service.SaveCashTransaction(body, transactionType)
	} else {
		transaction, err = c.service.SaveTransferTransaction(body)
	}

	if (err != nil) {
		return transaction, err
	}

	return transaction, nil
}

func getTransactionType(body map[string]interface{}) (string, error) {

	Type, ok := body["type"].(string)
	if ok != true {
		return Type, errors.New("Unable to find transaction type")
	}

	if (Type != "cash_deposit" && Type != "cash_withdraw" && Type != "account_transfer") {
		return Type, errors.New("Transaction type should be from [ cash_deposit , cash_withdraw , account_transfer ]")
	}
	return Type, nil
}
