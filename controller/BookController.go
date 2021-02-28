package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	. "../models"
	"github.com/gorilla/mux"
)

// GET All books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SelectAll())

}

// GET Single book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("single get")
}

//POST create
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var book Book
	err := decoder.Decode(&book)
	if err != nil {
		log.Fatal(err)
	} else {
		json.NewEncoder(w).Encode(book)
		CreateBookModel(book)
	}

}

//PUT update
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var book Book
	err := decoder.Decode(&book)
	if err != nil {
		log.Fatal(err)
	} else {
		json.NewEncoder(w).Encode(book)
		UpdateBookModel(params["id"], book)
	}
}

//DELETE delete
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	 id,_ := strconv.Atoi(params["id"])
	
	DeleteBookModel(id)
	//json.NewEncoder(w).Encode(params)
	//defer r.Body.Close()
}
