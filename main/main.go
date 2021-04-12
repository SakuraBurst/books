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
	router.Use(middleware.ContentTypeMiddleware)
	router.HandleFunc("/api/books", booksController.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", booksController.GetBook).Methods("GET")
	router.HandleFunc("/api/books/{id}", booksController.UpdateBook).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/books", booksController.AddBook).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/books/{id}", booksController.DeleteBook).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/registration", booksController.Registration).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/login", booksController.Login).Methods("POST", "OPTIONS")

	log.Fatal(http.ListenAndServe(":3585", router))
}
