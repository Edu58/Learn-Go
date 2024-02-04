package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	dsn := "postgres://ed:ed@localhost:5432/go_bookstore"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Could not connect to Database")
		panic(err)
	}
	Db = db
}

func GetDb() *gorm.DB {
	return Db
}
