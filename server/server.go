package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type Message struct {
	ID        int64
	Username  string
	Content   string
	Likes     int64
	Ys        int64
	CreatedAt string
}

type Messages []Message

// Converts a Message struct to JSON string.
func (a Message) ToJSON() string {
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return "{}"
	}
	return string(b)
}

// Converts a Messages slice to JSON string.
func (a Messages) ToJSON() string {
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return "[]"
	}
	return string(b)
}

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
		return nil, fmt.Errorf("Could not establish a connection to MySQL database at %s", cfg.Addr)
	}
	fmt.Println("Connected!")

	return db, nil
}

// Retrieves all messages from the database and returns a Messages slice.
func GetAllMessages(db *sql.DB) (Messages, error) {
	// An messages slice to hold data from returned rows.
	var messages Messages

	// Retrieve message data from database.
	rows, err := db.Query("SELECT * FROM messages ORDER BY created_at DESC")
	if err != nil {
		return nil, fmt.Errorf("getAllMessages %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var message Message

		if err := rows.Scan(&message.ID, &message.Username, &message.Content, &message.Likes, &message.Ys, &message.CreatedAt); err != nil {
			return nil, fmt.Errorf("getAllMessages %v", err)
		}
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getAllMessages %v", err)
	}

	return messages, nil
}

// Retrieves a message by id from the database and returns a Message struct.
func GetMessageById(db *sql.DB, id int) (Message, error) {
	var message Message

	row := db.QueryRow("SELECT * FROM messages WHERE id = ?", id)

	if err := row.Scan(&message.ID, &message.Username, &message.Content, &message.Likes, &message.Ys, &message.CreatedAt); err != nil {
		return message, fmt.Errorf("GetMessageById %v", err)
	}

	return message, nil
}

func main() {

	// Initialize database connection.
	db, err := InitializeDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	e := echo.New()

	// All messages
	e.GET("/messages", func(c echo.Context) error {
		messages, err := GetAllMessages(db)
		if err != nil {
			log.Fatal(err)
		}
		return c.String(http.StatusOK, messages.ToJSON())
	})

	// Message by id
	e.GET("/messages/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusOK, "[]")
		}
		message, err := GetMessageById(db, id)
		if err != nil {
			c.String(http.StatusOK, "[]")
		}
		return c.String(http.StatusOK, message.ToJSON())
	})
	e.Logger.Fatal(e.Start(":1323"))
}
