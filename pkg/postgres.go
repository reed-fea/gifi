package pkg

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

func initer() *gorm.DB {
	dbURL := "postgres://sgmnqtui:H94izUPgL0M8ue111kMF3gvJkeeIWbIN@tiny.db.elephantsql.com:5432/sgmnqtui"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(User{})
	db.Create(&User{Name: "里斯", Age: "29.5"})
	return db
}

func GetPostgres() {
	initer()
}
