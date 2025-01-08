package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func DBConn() {
	connString := "user=postgres dbname=webchat sslmode=disable password=postgres host=localhost"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalln("Could not connect to DB")
	}
	Db = db
}

func GetDB() *sql.DB {
	return Db
}
