package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/SakuraBurst/books.git/main/controllers"
	"github.com/SakuraBurst/books.git/main/controllers/middleware"
	"github.com/SakuraBurst/books.git/main/driver"
	"github.com/SakuraBurst/books.git/main/repository"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// TODO: EVENT EMMITER ДЛЯ SQL запросов
var database *sql.DB

func main() {
	router := mux.NewRouter()
	database = driver.ConnectDatabase("DB_URL")
	repo := repository.Repository{Database: database}
	booksController := controllers.Controler{Repository: repo}
	router.HandleFunc("/books", booksController.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", booksController.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", booksController.UpdateBook).Methods("PUT")
	router.HandleFunc("/books", booksController.AddBook).Methods("POST")
	router.HandleFunc("/books/{id}", booksController.DeleteBook).Methods("DELETE")
	router.HandleFunc("/registration", booksController.Registration).Methods("POST")
	router.Use(middleware.LoggingMiddleware)
	log.Fatal(http.ListenAndServe(":3585", router))
}
