package models

import (
	"github.com/Edu58/GoBookStore/src/config"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBook(id int) Book {
	var Result Book
	db.First(&Result, id)
	return Result
}

func (b *Book) UpdateBook(data Book) {
	db.Model(b).Clauses(clause.Returning{}).Updates(data)
}

func DeleteBook(id int) {
	db.Delete(&Book{}, id)
}
