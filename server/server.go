package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

func (a Message) ToJSON() string {
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return "{}"
	}
	return string(b)
}

func (a Messages) ToJSON() string {
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return "[]"
	}
	return string(b)
}

func getAllMessages(db *sql.DB) (Messages, error) {
	// An messages slice to hold data from returned rows.
	var messages Messages

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

func main() {
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
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	messages, err := getAllMessages(db)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.GET("/messages/all", func(c echo.Context) error {
		return c.String(http.StatusOK, messages.ToJSON())
	})

	e.GET("/messages/:id", func(c echo.Context) error {
		id := c.Param("id")
		for _, message := range messages {
			if id == fmt.Sprintf("%d", message.ID) {
				return c.String(http.StatusOK, message.ToJSON())
			}
		}
		return c.String(http.StatusOK, "[]")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
