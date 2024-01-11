package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// Initializes a MySQL database connection and returns a sql.DB object.
func InitializeDB() (*sql.DB, error) {
	var db *sql.DB

	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "123456",
		Net:    "tcp",
		Addr:   "127.0.0.1:7675",
		DBName: "chat",
	}
	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	pingErr := db.Ping()
	if pingErr != nil || err != nil {
		return nil, fmt.Errorf("could not establish a connection to MySQL database at %s", cfg.Addr)
	}
	fmt.Println("Connected!")

	return db, nil
}
