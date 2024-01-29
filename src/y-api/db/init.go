package db

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

// Information about the database connection.
// https://go.dev/doc/database/open-handle

// Initializes a MySQL database connection and returns a sql.DB object.
func getDB() (*sql.DB, error) {

	// err := godotenv.Load("example.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("Y_DB_SERVICE_SERVICE_HOST") + ":" + os.Getenv("Y_DB_SERVICE_SERVICE_PORT"),
		DBName: os.Getenv("MYSQL_DATABASE"),
	}
	// Get a database handle.
	sqlDB, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	var database = sqlDB

	// Set maximum number of connections.
	// database.SetMaxOpenConns(20)

	// Verify connection properties.
	if err := database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}
