package service

import (
	"apiass/db"
	"apiass/entity"
	"errors"
	"fmt"
)

type CustomerService interface {
	Save(customer entity.Customer) (entity.Customer, error)
	Update(customer entity.Customer) (entity.Customer, error)
	Find(Id int64) (entity.Customer, error)
	Delete(Id int64) error
}

type customerService struct {}

func NewCustomerService() CustomerService {
	return &customerService{}
}

func (service *customerService) Save(customer entity.Customer) (entity.Customer, error) {
	_ , err := db.Database.Model(&customer).Returning("*").Insert()
    if err != nil {
		fmt.Println(err)
		err = errors.New("Something went wrong")
		return customer, err
	}
	return customer, nil
}

func (service *customerService) Find(Id int64) (entity.Customer, error) {
	var customer entity.Customer
	err := db.Database.Model(&customer).Where("customer.id = ?", Id).Select()
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			err = errors.New("No customer found with given Id")
			return customer, err
		}
		fmt.Println(err)
		err = errors.New("Something went wrong")
		return customer, err
	}
	return customer, nil
}

func (service *customerService) Update(customer entity.Customer) (entity.Customer, error) {
	res, err := db.Database.Model(&customer).Returning("*").WherePK().UpdateNotZero()
	fmt.Println(res)
    if err != nil {
		if err.Error() == "pg: no rows in result set" {
			err = errors.New("No customer found with given Id")
			return customer, err
		}
		fmt.Println(err)
		err = errors.New("Something went wrong")
		return customer, err
    }

	fmt.Println(customer)
	// service.customers = append(service.customers, customer)
	return customer, nil
}

func (service *customerService) Delete(Id int64) error {
	var customer entity.Customer
	res, err := db.Database.Model(&customer).Where("id = ?", Id).Returning("*").Delete()
	fmt.Println(res, customer)
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			err = errors.New("No customer found with given Id")
			return err
		}
		fmt.Println(err)
		err = errors.New("Something went wrong")
		return err
	}
	return nil
}