package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s",
		dbUser,
		dbPass,
		dbHost,
		dbName,
	)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("failed to open db:", err)
	}

	// Test connection
	if err := DB.Ping(); err != nil {
		log.Fatal("failed to connect to db:", err)
	}

	// connection pool settings
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	log.Println("Database connected successfully")
}
