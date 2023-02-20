package controller

import (
	"apiass/entity"
	"apiass/service"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	bankService     service.BankService     = service.NewBankService()
	customerService service.CustomerService = service.NewCustomerService()
)

type AccountController interface {
	Find(ctx *gin.Context) (entity.Account, error)
	Save(ctx *gin.Context) (entity.Account, error)
	SaveJointAccount(ctx *gin.Context) (entity.Account, error)
	Delete(ctx *gin.Context) error
}

type accountController struct {
	service service.AccountService
}

func NewAccountController(service service.AccountService) AccountController {
	return &accountController{
		service: service,
	}
}

func (c *accountController) Find(ctx *gin.Context) (entity.Account, error) {
	var account entity.Account
	Number, err := strconv.ParseInt(ctx.Params.ByName("number"), 0, 64)
	if err != nil {
		return account, err
	}

	account, err = c.service.Find(Number)
	if err != nil {
		return account, err
	}
	return account, nil
}

func (c *accountController) Save(ctx *gin.Context) (entity.Account, error) {
	var account entity.Account

	var createAcc entity.CreateAccount
	err := ctx.BindJSON(&createAcc)
	if err != nil {
		return account, err
	}

	_, err = bankService.Find(createAcc.BankId)
	if err != nil {
		err = errors.New("bank_id: not found or invalid")
		return account, err
	}
	_, err = customerService.Find(createAcc.CustomerId)
	if err != nil {
		err = errors.New("customer_id: not found or invalid")
		return account, err
	}

	account, err = c.service.Save(createAcc.BankId, false)

	if err != nil {
		return account, err
	}

	err = c.service.MapAccountCustomer(account.Number, createAcc.CustomerId)

	return account, nil
}

func (c *accountController) SaveJointAccount(ctx *gin.Context) (entity.Account, error) {
	var account entity.Account

	var createAcc entity.CreateAccount
	err := ctx.BindJSON(&createAcc)
	if err != nil {
		return account, err
	}


	_, err = bankService.Find(createAcc.BankId)
	if err != nil {
		err = errors.New("bank_id: not found or invalid")
		return account, err
	}

	if len(createAcc.CustomerIds) != 2 || 
	createAcc.CustomerIds[0] == createAcc.CustomerIds[1] {
		err = errors.New("Exactly two DISTINCT customer_ids are required")
		return account, err
	}

	for _, customerId := range createAcc.CustomerIds {
		_, err = customerService.Find(customerId)
		if err != nil {
			err = errors.New("customer_id: not found or invalid")
			return account, err
		}
	}

	account, err = c.service.Save(createAcc.BankId, true)
	if err != nil {
		return account, err
	}

	for _, customerId := range createAcc.CustomerIds {
		err = c.service.MapAccountCustomer(account.Number, customerId)
		if (err != nil) {
			return account, err
		}
	}
	return account, nil
}

func (c *accountController) Delete(ctx *gin.Context) error {
	Number, err := strconv.ParseInt(ctx.Params.ByName("number"), 0, 64)
	if err != nil {
		return err
	}

	err = c.service.Delete(Number)
	if err != nil {
		return err
	}
	return nil
}
