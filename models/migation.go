package models

import (
	"theletter/models/user"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)



func CreateSchema(db *pg.DB) error {
	models := []interface{}{(*user.User)(nil)}
	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
			IfNotExists: true,
		})

		if err != nil {
			return err
		}
	}
	return nil
}