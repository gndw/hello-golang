package database

import (
	"fmt"
	"hello-golang/models"

	pg "github.com/go-pg/pg/v10"
	orm "github.com/go-pg/pg/v10/orm"
)

var Db *pg.DB

func Initialize() {
	Db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "example",
		Database: "mydb",
	})

	err := createSchema(Db)
	if err != nil {
		panic(err)
	}
}

func Shutdown() {
	fmt.Println("Closing Database Connection...")
	Db.Close()
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*models.User)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
