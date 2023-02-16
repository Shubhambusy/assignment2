package entity

type Customer struct {
	Id      int64	`json:"id"`
	Name    string	`json:"name"`
	Phone   int64	`json:"phone"`
	Address string	`json:"address"`
}
