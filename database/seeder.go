package database

import (
	"github.com/go-pg/pg/v9"
	"theletter/models/user"
)

func Seed(db *pg.DB) {
	user.CreateUser(db, "Nuno", "secret", "nunnomalex@gmail.com", "11-12-2019")
}