package service

import (
	"apiass/db"
	"apiass/entity"
	"errors"
	"fmt"
	"time"
)

type AccountService interface {
	Save(bankID int64, isJoint bool) (entity.Account, error)
	MapAccountCustomer(accountId int64, customerId int64) error
	Find(Number int64) (entity.Account, error)
	Delete(Number int64) error
}

type accountService struct{}

func NewAccountService() AccountService {
	return &accountService{}
}

func (service *accountService) Save(bankId int64, isJoint bool) (entity.Account, error) {

	var account entity.Account
	account.Bank_id = bankId
	account.OpeningDate = time.Now()
	account.JointAcccount = isJoint

	_, err := db.Database.Model(&account).Returning("*").Insert()
	if err != nil {
		fmt.Println(err)
		err = errors.New("Something went wrong")
		return account, err
	}

	return account, nil
}

func (service *accountService) Find(Number int64) (entity.Account, error) {
	var account entity.Account
	err := db.Database.Model(&account).Where("account.number = ?", Number).Select()
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			err = errors.New("No account found with given Account Number")
			return account, err
		}
		fmt.Println(err)
		err = errors.New("Something went wrong")
		return account, err
	}
	return account, nil
}

func (service *accountService) Delete(Number int64) error {
	var account entity.Account
	res, err := db.Database.Model(&account).Where("number = ?", Number).Returning("*").Delete()
	fmt.Println(res, account)
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			err = errors.New("No account found with given Account Number")
			return err
		}
		fmt.Println(err)
		err = errors.New("Something went wrong")
		return err
	}
	return nil
}

func (service *accountService)MapAccountCustomer(accountNumber int64, customerId int64) error {
	var accountCustomerMap entity.AccountCustomerMap
	accountCustomerMap.AccountNumber = accountNumber
	accountCustomerMap.CustomerId = customerId

	_, err := db.Database.Model(&accountCustomerMap).Returning("*").Insert()
	if err != nil {
		fmt.Println(err)
		err = errors.New("Something went wrong")
		return err
	}
	return nil
}