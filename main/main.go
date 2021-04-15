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
	spa := spaHandler{staticPath: "build", indexPath: "index.html"}
	router.HandleFunc("/api/books", booksController.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/api/books/{id}", booksController.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/api/books/{id}", booksController.UpdateBook).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/api/books", booksController.AddBook).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/books/{id}", booksController.DeleteBook).Methods(http.MethodDelete, http.MethodOptions)
	router.HandleFunc("/api/registration", booksController.Registration).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/login", booksController.Login).Methods(http.MethodPost, http.MethodOptions)
	router.PathPrefix("/").Handler(spa)

	log.Fatal(http.ListenAndServe(":3584", router))
}
