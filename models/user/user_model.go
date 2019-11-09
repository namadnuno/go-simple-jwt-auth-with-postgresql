package user

import (
	"github.com/go-pg/pg/v9"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	Id        int64
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func GetUserByLogin(db *pg.DB, email string, plainPassword string) *User {
	user := new(User)
	err := db.Model(user).Where("email = ?", email).Select()

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return user
}

func GetAllUsers(db *pg.DB) []User {
	var users []User
	err := db.Model(&users).Select()

	if err != nil {
		panic(err)
	}

	return users;
}

func CreateUser(db *pg.DB, name string, plainPassword string, email string, createdAt string) *User {
	password, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	u := &User{
		Name: name,
		Password: string(password),
		Email: email,
		CreatedAt: createdAt,
	}

	err = db.Insert(u)

	if err != nil {
		panic(err)
	}

	return u;
}