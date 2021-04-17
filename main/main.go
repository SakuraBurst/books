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

	database = driver.ConnectDatabase("DB_URL")
	repo := repository.Repository{Database: database}
	booksController := controllers.Controler{Repository: repo}
	router := router(booksController)

	log.Fatal(http.ListenAndServe(":3584", router))
}

func router(booksController controllers.Controler) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.ContentTypeMiddleware)
	spa := spaHandler{staticPath: "build", indexPath: "index.html"}
	router.HandleFunc("/api/registration", booksController.Registration).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/login", booksController.Login).Methods(http.MethodPost, http.MethodOptions)
	protected := router.PathPrefix("/api/v2").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/books", booksController.GetBooks).Methods(http.MethodGet)
	protected.HandleFunc("/books/{id}", booksController.GetBook).Methods(http.MethodGet)
	protected.HandleFunc("/books/{id}", booksController.UpdateBook).Methods(http.MethodPut, http.MethodOptions)
	protected.HandleFunc("/books", booksController.AddBook).Methods(http.MethodPost, http.MethodOptions)
	protected.HandleFunc("/books/{id}", booksController.DeleteBook).Methods(http.MethodDelete, http.MethodOptions)
	protected.HandleFunc("/user", booksController.UpdateUser).Methods(http.MethodGet)

	router.PathPrefix("/").Handler(spa)
	return router
}
