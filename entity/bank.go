package entity

type Bank struct {
	Id		int64	`json:"id"`
	Name	string	`json:"name" pg:",unique,notnull"`
	Address	string	`json:"address"`
}