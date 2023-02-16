package db

import (
	"apiass/entity"
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var Database *pg.DB

func InitDB () {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "123",
		Database: "mydb",
	})

	ctx := context.Background()
	
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}
	fmt.Println("connected to database")

	err := createSchema(db)
	if err != nil {
		panic(err)
	}
	Database = db
}

// createSchema creates database schema
func createSchema(db *pg.DB) error {

    models := []interface{}{
        (*entity.Bank)(nil),
		(*entity.Account)(nil),
		(*entity.Customer)(nil),
		(*entity.Transaction)(nil),
    }

    for _, model := range models {
        err := db.Model(model).CreateTable(&orm.CreateTableOptions{
            IfNotExists: true,			
        })
        if err != nil {
            return err
        }
		fmt.Println("table created")
    }

    return nil
}