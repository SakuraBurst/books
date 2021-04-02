package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func ConnectDatabase(envKey string) *sql.DB {
	dbUrl := getEnv(envKey)
	db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("ну значит у тебя нет енвх што поделать то")
	}
	return os.Getenv(key)
}
