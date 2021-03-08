package routes

import (
	"log"
	"net/http"

	. "github.com/glburak/golangRestApiMysql/controller"
	"github.com/gorilla/mux"
)

func Routers() {

	r := mux.NewRouter()

	r.HandleFunc("/api/books", GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/api/books", CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", DeleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))

}
