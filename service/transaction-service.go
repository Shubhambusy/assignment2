package service

import (
	"apiass/db"
	"apiass/entity"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TransactionService interface {
	Save(transaction entity.Transaction) (entity.Transaction, error)
	SaveCashTransaction(transactionData map[string]interface{},
		transactionType string) (entity.Transaction, error)
	SaveTransferTransaction(transactionData map[string]interface{}) (entity.Transaction, error)
	Find(Id uuid.UUID) (entity.Transaction, error)
}

type transactionService struct{}

func NewTransactionService() TransactionService {
	return &transactionService{}
}

func (service *transactionService) SaveCashTransaction(
	transactionData map[string]interface{},
	transactionType string) (entity.Transaction, error) {
	var transaction entity.Transaction

	transaction.Type = transactionType
	transaction.Timestamp = time.Now()
	amount, ok := transactionData["amount"].(float64)
	if ok == false {
		return transaction, errors.New("Amount could not be fetched")
	}
	transaction.Amount = amount

	float_account_number, ok := transactionData["account_number"].(float64)
	account_number := int64(float_account_number) // json does not support int, it onlu considerrs float
	if ok == false {
		return transaction, errors.New("Account Number could not be fetched")
	}
	if transactionType == "cash_deposit" {
		transaction.RecipientAcount = account_number
		transaction.SenderAccount = 0
	} else if transactionType == "cash_withdraw" {
		transaction.SenderAccount = account_number
		transaction.RecipientAcount = 0
	} else {
		panic("Transection Type was not handled properly")
	}

	transaction, err := service.Save(transaction)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (service *transactionService) SaveTransferTransaction(
	transactionData map[string]interface{}) (entity.Transaction, error) {
	var transaction entity.Transaction
	transaction.Type = "account_transfer"

	amount, ok := transactionData["amount"].(float64)
	if ok == false {
		return transaction, errors.New("Amount could not be fetched")
	}
	transaction.Amount = amount

	float_sender_account, ok := transactionData["sender_account"].(float64)
	sender_account := int64(float_sender_account) // json does not support int, it onlu considerrs float
	if ok == false {
		return transaction, errors.New("Sender Account Number could not be fetched")
	}
	float_recipient_account, ok := transactionData["recipient_account"].(float64)
	recipient_account := int64(float_recipient_account) // json does not support int, it onlu considerrs float
	if ok == false {
		return transaction, errors.New("Recipient Account Number could not be fetched")
	}
	transaction.SenderAccount = sender_account
	transaction.RecipientAcount = recipient_account

	transaction.Timestamp = time.Now()

	transaction, err := service.Save(transaction)
	if err != nil {
		return transaction, err
	}
	
	return transaction, nil
}

func (service *transactionService) Find(Id uuid.UUID) (entity.Transaction, error) {
	var transaction entity.Transaction
	err := db.Database.Model(&transaction).Where("transaction.id = ?", Id).Select()
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			err = errors.New("No transaction found with given Id")
			return transaction, err
		}
		panic(err)
	}
	return transaction, nil
}

func (service *transactionService) Save(
	transaction entity.Transaction) (entity.Transaction, error) {

	transaction.Id = uuid.New()

	if (transaction.Amount < 0) {
		err := errors.New("Amount can not be negetive")
		return transaction, err
	}
	err := updateAccountBalance(transaction.SenderAccount,
		transaction.RecipientAcount,
		transaction.Amount)
	if err != nil {
		return transaction, err
	}

	_, err = db.Database.Model(&transaction).Returning("*").Insert()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return transaction, nil
}

func updateAccountBalance(
	SenderAccountNumber int64,
	RecipientAccountNumber int64,
	Amount float64,
) error {

	accountService := NewAccountService()
	var err error
	var senderAccount entity.Account
	var recipientAccount entity.Account
	if SenderAccountNumber != 0 {
		senderAccount, err = accountService.Find(SenderAccountNumber)
		if err != nil {
			return errors.New("Debetting Account could not be validated")
		}
		if senderAccount.Balance < Amount {
			return errors.New("Debetting Account balance is not sufficient")
		}
		senderAccount.Balance = senderAccount.Balance - Amount
	}

	if RecipientAccountNumber != 0 {
		fmt.Println(RecipientAccountNumber)
		recipientAccount, err = accountService.Find(RecipientAccountNumber)
		if err != nil {
			return errors.New("Creditting Account could not be validated")
		}
		recipientAccount.Balance = recipientAccount.Balance + Amount
	}

	fmt.Println("sender acc --", senderAccount)
	fmt.Println("Recipient acc. -- ", recipientAccount)
	if SenderAccountNumber != 0 {
		_, err := db.Database.Model(&senderAccount).WherePK().Update()
		if err != nil {
			panic(err)
		}
	}


	if RecipientAccountNumber != 0 {
		_, err := db.Database.Model(&recipientAccount).WherePK().Update()
		if err != nil {
			panic(err)
		}
	}

	return nil
}
