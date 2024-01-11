package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

type Message struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Content   string `json:"content"`
	Likes     int64  `json:"likes"`
	Ys        int64  `json:"ys"`
	CreatedAt string `json:"created_at"`
}

type Messages []Message

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

// Add a message to the database and return the message struct
func AddMessage(db *sql.DB, message Message) (Message, error) {
	stmt, err := db.Prepare("INSERT INTO messages (username, content, likes, ys) VALUES (?, ?, ?, ?)")
	if err != nil {
		return message, err
	}
	res, err := stmt.Exec(message.Username, message.Content, message.Likes, message.Likes)
	if err != nil {
		return message, err
	}
	// Return success response
	id, err := res.LastInsertId()
	if err != nil {
		return message, err
	}
	message.ID = id
	return message, nil
}

func main() {
	// Initialize database connection.
	db, err := InitializeDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Echo instance
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// https://echo.labstack.com/docs/middleware/cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// All messages
	e.GET("/messages", func(c echo.Context) error {
		messages, err := GetAllMessages(db)

		if err != nil {
			return c.JSON(http.StatusOK, Messages{})
		}
		return c.JSON(http.StatusOK, messages)
	})

	// Message by id
	e.GET("/messages/:id", func(c echo.Context) error {
		// Convert id param to int
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusOK, "Invalid id")
		}
		// Retrieve message by id
		message, err := GetMessageById(db, id)
		if err != nil {
			c.JSON(http.StatusOK, Message{})
		}
		return c.JSON(http.StatusOK, message)
	})

	// Add a message
	e.POST("/messages", func(c echo.Context) error {
		// example body: { "username": "test", "content": "test", "likes": 0, "ys": 0 }
		// Bind message data from request body to Message struct
		var message Message
		if err := c.Bind(&message); err != nil {
			return err
		}
		// Add message to database
		message, err = AddMessage(db, message)
		if err != nil {
			return err
		}
		// Return success response
		return c.JSON(http.StatusCreated, &message)
	})

	// Initialize router
	e.Logger.Fatal(e.Start(":1323"))
}
