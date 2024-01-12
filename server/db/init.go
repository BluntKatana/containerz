package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Information about the database connection.
// https://go.dev/doc/database/open-handle

// A sql.DB object manages a pool of database connections.
var database *sql.DB

// Initializes a MySQL database connection and returns a sql.DB object.
func init() {
	// Load .env file
	err := godotenv.Load("example.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_ROOT_USERNAME"),
		Passwd: os.Getenv("MYSQL_ROOT_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT"),
		DBName: os.Getenv("MYSQL_DATABASE"),
	}
	// Get a database handle.
	sqlDB, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	database = sqlDB

	// Set maximum number of connections.
	// database.SetMaxOpenConns(20)

	// Verify connection properties.
	if err := database.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Init - Successfully connected!")
}

func Ping() error {
	if err := database.Ping(); err != nil {
		return err
	}
	fmt.Println("Ping - Successfully connected!")
	return nil
}
