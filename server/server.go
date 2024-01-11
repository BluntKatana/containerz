package main

import (
	"net/http"
	"strconv"

	db "db"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

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

		// Initialize database connection.
		database, err := db.InitializeDB()
		if err != nil {
			return echo.NewHTTPError(http.StatusServiceUnavailable, "something went wrong processing your request")
		}
		defer database.Close()

		messages, err := db.GetAllMessages(database)
		if err != nil {
			return c.JSON(http.StatusOK, db.Messages{})
		}
		return c.JSON(http.StatusOK, messages)
	})

	// Message by id
	e.GET("/messages/:id", func(c echo.Context) error {

		// Initialize database connection.
		database, err := db.InitializeDB()
		if err != nil {
			return echo.NewHTTPError(http.StatusServiceUnavailable, "something went wrong processing your request")
		}
		defer database.Close()

		// Convert id param to int
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid message id (must be a number)")
		}
		// Retrieve message by id
		message, err := db.GetMessageById(database, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "message not found")
		}
		return c.JSON(http.StatusOK, message)
	})

	// Add a message
	e.POST("/messages", func(c echo.Context) error {

		// Initialize database connection.
		database, err := db.InitializeDB()
		if err != nil {
			return echo.NewHTTPError(http.StatusServiceUnavailable, "something went wrong processing your request")
		}
		defer database.Close()

		// example body: { "username": "test", "content": "test", "likes": 0, "ys": 0 }
		// Bind message data from request body to Message struct
		var message db.Message
		if err := c.Bind(&message); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid message data")
		}

		// Check length requirements
		if len(message.Content) > 155 {
			return echo.NewHTTPError(http.StatusBadRequest, "message content too long")
		}

		// Add message to database
		message, err = db.AddMessage(database, message)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "could not process request")
		}
		// Return success response
		return c.JSON(http.StatusCreated, &message)
	})

	// Update a message: add a like
	e.PUT("/messages/:id/like", func(c echo.Context) error {
		// Initialize database connection.
		database, err := db.InitializeDB()
		if err != nil {
			return echo.NewHTTPError(http.StatusServiceUnavailable, "something went wrong processing your request")
		}
		defer database.Close()

		// Convert id param to int
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid message id (must be a number)")
		}
		// Retrieve message by id
		message, err := db.GetMessageById(database, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "message not found")
		}

		message.Likes = message.Likes + 1

		// Update message in database
		message, err = db.UpdateMessage(database, message)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "could not process request")
		}
		// Return success response
		return c.JSON(http.StatusCreated, &message)
	})

	// Update a message: remove a like
	e.PUT("/messages/:id/unlike", func(c echo.Context) error {
		// Initialize database connection.
		database, err := db.InitializeDB()
		if err != nil {
			return echo.NewHTTPError(http.StatusServiceUnavailable, "something went wrong processing your request")
		}
		defer database.Close()

		// Convert id param to int
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid message id (must be a number)")
		}
		// Retrieve message by id
		message, err := db.GetMessageById(database, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "message not found")
		}

		message.Likes = message.Likes - 1

		// Update message in database
		message, err = db.UpdateMessage(database, message)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "could not process request")
		}
		// Return success response
		return c.JSON(http.StatusCreated, &message)
	})

	// Update a message: add a y
	e.PUT("/messages/:id/y", func(c echo.Context) error {

		// Initialize database connection.
		database, err := db.InitializeDB()
		if err != nil {
			return echo.NewHTTPError(http.StatusServiceUnavailable, "something went wrong processing your request")
		}
		defer database.Close()

		// Convert id param to int
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid message id (must be a number)")
		}
		// Retrieve message by id
		message, err := db.GetMessageById(database, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "message not found")
		}

		message.Ys = message.Ys + 1

		// Update message in database
		message, err = db.UpdateMessage(database, message)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "could not process request")
		}
		// Return success response
		return c.JSON(http.StatusCreated, &message)
	})

	// Initialize router
	e.Logger.Fatal(e.Start(":1323"))
}
