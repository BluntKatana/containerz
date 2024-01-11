package main

import (
	"net/http"
	"strconv"

	db "db"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func getAllMessages(c echo.Context) error {
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
}

func getMessage(c echo.Context) error {
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
}

func createMessage(c echo.Context) error {
	// Initialize database connection.
	database, err := db.InitializeDB()
	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "something went wrong processing your request")
	}
	defer database.Close()

	// Bind message data from request body to Message struct
	var message db.Message
	if err := c.Bind(&message); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid message data")
	}
	// Ensure message content is no longer than 255 characters
	if len(message.Content) > 255 {
		return echo.NewHTTPError(http.StatusBadRequest, "message content must be less than 255 characters")
	}

	// Add message to database
	message, err = db.AddMessage(database, message)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not process request")
	}
	// Return success response
	return c.JSON(http.StatusCreated, &message)
}

func likeMessage(c echo.Context) error {
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
}

func unlikeMessage(c echo.Context) error {
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
}

func addyMessage(c echo.Context) error {
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
}

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

	// Routes
	e.GET("/messages", getAllMessages)
	e.GET("/messages/:id", getMessage)
	e.POST("/messages", createMessage)
	e.PUT("/messages/:id/like", likeMessage)
	e.PUT("/messages/:id/unlike", unlikeMessage)
	e.PUT("/messages/:id/y", addyMessage)

	// Initialize router
	e.Logger.Fatal(e.Start(":1323"))
}
