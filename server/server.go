package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	db "db"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Initialize database connection.
	database, err := db.InitializeDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// https://echo.labstack.com/docs/middleware/cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// All messages
	e.GET("/messages", func(c echo.Context) error {
		messages, err := db.GetAllMessages(database)

		if err != nil {
			return c.JSON(http.StatusOK, db.Messages{})
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
		message, err := db.GetMessageById(database, id)
		if err != nil {
			// return not found error
			return c.JSON(http.StatusOK, db.Message{})
		}
		return c.JSON(http.StatusOK, message)
	})

	// Add a message
	e.POST("/messages", func(c echo.Context) error {
		// example body: { "username": "test", "content": "test", "likes": 0, "ys": 0 }
		// Bind message data from request body to Message struct
		var message db.Message
		if err := c.Bind(&message); err != nil {
			return err
		}
		// Add message to database
		message, err = db.AddMessage(database, message)
		if err != nil {
			return err
		}
		// Return success response
		return c.JSON(http.StatusCreated, &message)
	})

	e.PUT("/messages/:id/like", func(c echo.Context) error {
		// Convert id param to int
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusOK, "Invalid id")
		}
		// Retrieve message by id
		message, err := db.GetMessageById(database, id)

		if err != nil {
			fmt.Println("err", err)
			// return not found error
			return c.JSON(http.StatusOK, db.Message{})
		}

		message.Likes = message.Likes + 1

		// Update message in database
		message, err = db.UpdateMessage(database, message)
		if err != nil {
			return err
		}
		// Return success response
		return c.JSON(http.StatusCreated, &message)
	})

	// Update a message its likes, ys, and/or content
	e.PUT("/messages/:id", func(c echo.Context) error {
		// example body: { "content": "test", "likes": 0, "ys": 0 }
		// Bind message data from request body to Message struct
		// Convert id param to int
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusOK, "Invalid id")
		}
		// Retrieve message by id
		message, err := db.GetMessageById(database, id)

		if err != nil {
			fmt.Println("err", err)
			// return not found error
			return c.JSON(http.StatusOK, db.Message{})
		}

		// Retrieve ys, likes and/or content from the request body
		// and update the message struct if they are present
		likes := c.FormValue("likes")
		if likes != "" {
			fmt.Println("likes", likes)
			likes, err := strconv.Atoi(likes)
			if err != nil {
				return c.String(http.StatusOK, "Invalid likes")
			}
			message.Likes = int64(likes)
		}

		ys := c.FormValue("ys")
		if ys != "" {
			ys, err := strconv.Atoi(ys)
			if err != nil {
				return c.String(http.StatusOK, "Invalid ys")
			}
			message.Ys = int64(ys)
		}

		content := c.FormValue("content")
		if content != "" {
			message.Content = content
		}

		// Update message in database
		message, err = db.UpdateMessage(database, message)
		if err != nil {
			return err
		}
		// Return success response
		return c.JSON(http.StatusCreated, &message)
	})

	// Initialize router
	e.Logger.Fatal(e.Start(":1323"))
}
