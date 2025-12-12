// database connection
package utils

import (
    "fmt"
    "log"
    "os"

    "github.com/jmoiron/sqlx"
    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

// Package-level variables
var (
    host     string
    port     string
    user     string
    password string
    dbname   string
)

// init runs automatically before main
func init() {
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: .env file not found")
    }

    host = os.Getenv("DB_HOST")
    port = os.Getenv("DB_PORT")
    user = os.Getenv("DB_USER")
    password = os.Getenv("DB_PASSWORD")
    dbname = os.Getenv("DB_NAME")
}

func GetConnection() *sqlx.DB {
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sqlx.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    log.Println("DB Connection established...")
    return db
}