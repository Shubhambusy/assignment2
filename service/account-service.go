package service

import (
	"apiass/db"
	"apiass/entity"
	"errors"
	"fmt"
)

type AccountService interface {
	Save(account entity.Account) entity.Account
	Update(account entity.Account) (entity.Account, error)
	Find(Id int64) (entity.Account, error)
	Delete(Id int64) error
}

type accountService struct{}

func NewAccountService() AccountService {
	return &accountService{}
}

func (service *accountService) Save(account entity.Account) entity.Account {
	_, err := db.Database.Model(&account).Returning("*").Insert()
	if err != nil {
		panic(err)
	}

	fmt.Println(account)
	// service.accounts = append(service.accounts, account)
	return account
}

func (service *accountService) Find(Id int64) (entity.Account, error) {
	var account entity.Account
	err := db.Database.Model(&account).Where("account.id = ?", Id).Select()
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			err = errors.New("No account found with given Id")
			return account, err
		}
		panic(err)
	}
	return account, nil
}

func (service *accountService) Update(account entity.Account) (entity.Account, error) {
	res, err := db.Database.Model(&account).Returning("*").WherePK().Update()
	fmt.Println(res)
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			err = errors.New("No account found with given Id")
			return account, err
		}
		panic(err)
	}

	fmt.Println(account)
	// service.accounts = append(service.accounts, account)
	return account, nil
}

func (service *accountService) Delete(Id int64) error {
	var account entity.Account
	res, err := db.Database.Model(&account).Where("id = ?", Id).Returning("*").Delete()
	fmt.Println(res, account)
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			err = errors.New("No account found with given Id")
			return err
		}
		panic(err)
	}
	return nil
}
