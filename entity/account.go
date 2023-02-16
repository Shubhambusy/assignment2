package entity

import "time"

type AccountCustomerMap struct {
	Id        int64
	AccountId int64
	Customer  int64
}

type Account struct {
	Number      int64 `pg:",pk"`
	Balance     int64
	BankId      int64
	OpeningDate time.Time
}
