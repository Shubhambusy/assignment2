package controller

import (
	"apiass/entity"
	"apiass/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	Find(ctx *gin.Context) (entity.Customer, error)
	Save(ctx *gin.Context) (entity.Customer, error)
	Update(ctx *gin.Context) (entity.Customer, error)
	Delete(ctx *gin.Context) error
}

type customerController struct {
	service service.CustomerService
}

func NewCustomerController(service service.CustomerService) CustomerController {
	return &customerController {
		service: service,
	}
}

func (c *customerController)Find(ctx *gin.Context) (entity.Customer, error) {
	var customer entity.Customer
	Id, err := strconv.ParseInt(ctx.Params.ByName("id"), 0, 64)
	if err != nil {
		return customer, err
	}
	
	customer, err = c.service.Find(Id)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (c *customerController) Save(ctx *gin.Context) (entity.Customer, error) {
	var customer entity.Customer
	err := ctx.BindJSON(&customer)
	if (err != nil) {
		return customer, err
	}
	customer, err = c.service.Save(customer)
	if (err != nil) {
		return  customer, err
	}
	return customer, nil
}

func (c *customerController) Update(ctx *gin.Context) (entity.Customer, error) {
	var customer entity.Customer
	Id, err := strconv.ParseInt(ctx.Params.ByName("id"), 0, 64)
	err = ctx.BindJSON(&customer)
	if err != nil {
		return customer, err
	}
	customer.Id = Id
	customer, err = c.service.Update(customer)
	return customer, err
}

func (c *customerController) Delete(ctx *gin.Context) error {
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