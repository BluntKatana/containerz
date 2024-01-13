package handler

import (
	"db"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// GetAllMessages is func get all messages
func GetAllMessages(c echo.Context) error {
	messages, err := db.GetAllMessages()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, messages)
}

// GetMessage is func get one message by
func GetMessage(c echo.Context) error {
	// Convert id param to int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	// Retrieve message by id
	message, err := db.GetMessageById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, message)
}

func CreateMessage(c echo.Context) error {
	// Bind message data from request body to Message struct
	var objRequest db.Message
	if err := c.Bind(&objRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "CRASH")
	}

	// Ensure message content is no longer than 255 characters
	if len(objRequest.Content) > 255 {
		return echo.NewHTTPError(http.StatusBadRequest, "message content must be less than 255 characters")
	}

	// Add message to database
	data, err := db.AddMessage(&objRequest)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	// Return success response
	return c.JSON(http.StatusCreated, data)
}

func LikeMessage(c echo.Context) error {
	// Convert id param to int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Retrieve message by id
	message, err := db.GetMessageById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	message.Likes = message.Likes + 1

	// Update message in database
	data, err := db.UpdateMessage(message)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	// Return success response
	return c.JSON(http.StatusOK, data)
}

func UnlikeMessage(c echo.Context) error {
	// Convert id param to int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	// Retrieve message by id
	message, err := db.GetMessageById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	message.Likes = message.Likes - 1

	// Update message in database
	data, err := db.UpdateMessage(message)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	// Return success response
	return c.JSON(http.StatusOK, data)
}

func AddyMessage(c echo.Context) error {
	// Convert id param to int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	// Retrieve message by id
	message, err := db.GetMessageById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	message.Ys = message.Ys + 1

	// Update message in database
	data, err := db.UpdateMessage(message)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	// Return success response
	return c.JSON(http.StatusOK, data)
}
