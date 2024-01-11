package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

// Information about the database connection.
// https://go.dev/doc/database/open-handle

// A sql.DB object manages a pool of database connections.
var database *sql.DB

// Initializes a MySQL database connection and returns a sql.DB object.
func init() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "123456",
		Net:    "tcp",
		Addr:   "127.0.0.1:7675",
		DBName: "chat",
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
