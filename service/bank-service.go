package service

import (
	"apiass/db"
	"apiass/entity"
	"errors"
	"fmt"
)

type BankService interface {
	Save(bank entity.Bank) (entity.Bank, error)
	Update(bank entity.Bank) (entity.Bank, error)
	Find(Id int64) (entity.Bank, error)
	Delete(Id int64) error
}

type bankService struct{}

func NewBankService() BankService {
	return &bankService{}
}

func (service *bankService) Save(bank entity.Bank) (entity.Bank, error) {
	_, err := db.Database.Model(&bank).Returning("*").Insert()
	if err != nil {
		if err.Error()[:12] == "ERROR #23505" {
			err = errors.New("Bank name Already exists")
			return bank, err
		}
		panic(err)
	}

	fmt.Println(bank)
	return bank, nil
}

func (service *bankService) Find(Id int64) (entity.Bank, error) {
	var bank entity.Bank
	err := db.Database.Model(&bank).Where("bank.id = ?", Id).Select()
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			err = errors.New("No bank found with given Id")
			return bank, err
		}
		panic(err)
	}
	return bank, nil
}

func (service *bankService) Update(bank entity.Bank) (entity.Bank, error) {
	res, err := db.Database.Model(&bank).Returning("*").WherePK().UpdateNotZero()
	fmt.Println(res)
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			err = errors.New("No bank found with given Id")
			return bank, err
		}
		panic(err)
	}

	fmt.Println(bank)
	return bank, nil
}

func (service *bankService) Delete(Id int64) error {
	var bank entity.Bank
	res, err := db.Database.Model(&bank).Where("id = ?", Id).Returning("*").Delete()
	fmt.Println(res, bank)
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			err = errors.New("No bank found with given Id")
			return err
		}
		panic(err)
	}
	return nil
}
