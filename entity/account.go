package entity

import "time"

type CreateAccount struct {
	BankId      int64   `json:"bank_id"`
	CustomerId  int64   `json:"customer_id"`
	CustomerIds []int64 `json:"customer_ids"`
}

type AccountCustomerMap struct {
	Id            int64
	AccountNumber int64
	CustomerId    int64
}

type Account struct {
	Number        int64 `pg:",pk"`
	Balance       float64 `pg:"default:0"`
	Bank_id       int64 `json:"bank_id"`
	Bank          *Bank `pg:"rel:has-one"`
	OpeningDate   time.Time
	JointAcccount bool `pg:"default:false"`
}
