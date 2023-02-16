package controller

import (
	"apiass/entity"
	"apiass/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BankController interface {
	Find(ctx *gin.Context) (entity.Bank, error)
	Save(ctx *gin.Context) (entity.Bank, error)
	Update(ctx *gin.Context) (entity.Bank, error)
	Delete(ctx *gin.Context) error
}

type bankController struct {
	service service.BankService
}

func NewBankController(service service.BankService) BankController {
	return &bankController{
		service: service,
	}
}

func (c *bankController) Find(ctx *gin.Context) (entity.Bank, error) {
	var bank entity.Bank
	Id, err := strconv.ParseInt(ctx.Params.ByName("id"), 0, 64)
	if err != nil {
		return bank, err
	}

	bank, err = c.service.Find(Id)
	if err != nil {
		return bank, err
	}
	return bank, nil
}

func (c *bankController) Save(ctx *gin.Context) (entity.Bank, error) {

	var bank entity.Bank
	err := ctx.BindJSON(&bank)
	if err != nil {
		return bank, err
	}
	bank, err = c.service.Save(bank)
	if err != nil {
		return bank, err
	}
	return bank, nil
}

func (c *bankController) Update(ctx *gin.Context) (entity.Bank, error) {
	var bank entity.Bank
	Id, err := strconv.ParseInt(ctx.Params.ByName("id"), 0, 64)

	err = ctx.BindJSON(&bank)
	if err != nil {
		return bank, err
	}
	bank.Id = Id
	bank, err = c.service.Update(bank)
	return bank, err
}

func (c *bankController) Delete(ctx *gin.Context) error {
	Id, err := strconv.ParseInt(ctx.Params.ByName("id"), 0, 64)
	if err != nil {
		return err
	}

	err = c.service.Delete(Id)
	if err != nil {
		return err
	}
	return nil
}
