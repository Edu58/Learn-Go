package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Edu58/GoBookStore/src/models"
	"github.com/Edu58/GoBookStore/src/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetBooks()
	res, err := json.Marshal(books)
	if err != nil {
		log.Fatalln("Could get books", err)
		w.WriteHeader(400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["id"]
	stringBookId, err := strconv.Atoi(bookId)
	if err != nil {
		errMessage := "Could not find Book ID or Book ID is not a string int"
		log.Println(errMessage, err)
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errMessage)
	} else {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		book := models.GetBook(stringBookId)
		json.NewEncoder(w).Encode(book)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	book := createBook.CreateBook()
	json.NewEncoder(w).Encode(&book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["id"]
	stringBookId, err := strconv.Atoi(bookId)
	if err != nil {
		errMessage := "Could not find Book ID or Book ID is not a string int"
		log.Println(errMessage, err)
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errMessage)
	} else {
		w.WriteHeader(204)
		w.Header().Set("Content-Type", "application/json")
		models.DeleteBook(stringBookId)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["id"]
	stringBookId, _ := strconv.Atoi(bookId)
	var updateBookData = &models.Book{}
	utils.ParseBody(r, updateBookData)
	book := models.GetBook(stringBookId)
	book.UpdateBook(*updateBookData)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
