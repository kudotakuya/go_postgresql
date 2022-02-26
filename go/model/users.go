package model

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type User struct {
	Name string `query:"name"`
	Age  int    `query:"age"`
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=mydb password=password sslmode=disable")

	if err != nil {
		log.Fatalln("接続失敗", err)
	}

	return db
}

func GetUsers(query *User) []user {
	db := connectDB()
	var users []user
	db.Debug().Where(query).Find(&users)

	defer db.Close()
	return users
}

func CreateUser(query *User) {
	db := connectDB()
	db.Create(query)

	defer db.Close()
}
