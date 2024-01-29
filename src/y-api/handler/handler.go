package handler

import (
	"db"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// GetAllMessages is func get all messages
func GetAllMessages(c echo.Context) error {
	messages, err := db.GetAllMessages()

	// Return error response if there is an error retrieving messages
	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, SERVICE_ERROR)
	}

	return c.JSON(http.StatusOK, messages)
}

// GetMessage is func get one message by
func GetMessage(c echo.Context) error {
	// Convert id param to int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NO_INT_ERROR)
	}

	// Retrieve message by id
	message, err := db.GetMessageById(id)
	if err != nil {
		if err.Error() == "connection_error" {
			return echo.NewHTTPError(http.StatusServiceUnavailable, SERVICE_ERROR)
		} else {
			return echo.NewHTTPError(http.StatusNotFound, NOT_FOUND_ERROR)
		}
	}

	return c.JSON(http.StatusOK, message)
}

func CreateMessage(c echo.Context) error {
	// Bind message data from request body to Message struct
	var objRequest db.Message
	if err := c.Bind(&objRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, BAD_REQUEST_ERROR)
	}

	// Ensure message content is no longer than 255 characters
	if len(objRequest.Content) > 255 {
		return echo.NewHTTPError(http.StatusBadRequest, MESSAGE_TOO_LONG_ERROR)
	}

	// Add message to database
	data, err := db.AddMessage(&objRequest)
	if err != nil {
		if err.Error() == "connection_error" {
			return echo.NewHTTPError(http.StatusServiceUnavailable, SERVICE_ERROR)
		} else {
			return echo.NewHTTPError(http.StatusNotFound, NOT_FOUND_ERROR)
		}
	}

	// Return success response
	return c.JSON(http.StatusCreated, data)
}

func LikeMessage(c echo.Context) error {
	// Convert id param to int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NO_INT_ERROR)
	}

	// Retrieve message by id
	message, err := db.GetMessageById(id)
	if err != nil {
		if err.Error() == "connection_error" {
			return echo.NewHTTPError(http.StatusServiceUnavailable, SERVICE_ERROR)
		} else {
			return echo.NewHTTPError(http.StatusNotFound, NOT_FOUND_ERROR)
		}
	}

	message.Likes = message.Likes + 1

	// Update message in database
	data, err := db.UpdateMessage(message)
	if err != nil {
		if err.Error() == "connection_error" {
			return echo.NewHTTPError(http.StatusServiceUnavailable, SERVICE_ERROR)
		} else {
			return echo.NewHTTPError(http.StatusNotFound, NOT_FOUND_ERROR)
		}
	}
	// Return success response
	return c.JSON(http.StatusOK, data)
}

func UnlikeMessage(c echo.Context) error {
	// Convert id param to int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NO_INT_ERROR)
	}
	// Retrieve message by id
	message, err := db.GetMessageById(id)
	if err != nil {
		if err.Error() == "connection_error" {
			return echo.NewHTTPError(http.StatusServiceUnavailable, SERVICE_ERROR)
		} else {
			return echo.NewHTTPError(http.StatusNotFound, NOT_FOUND_ERROR)
		}
	}

	message.Likes = message.Likes - 1

	// Update message in database
	data, err := db.UpdateMessage(message)
	if err != nil {
		if err.Error() == "connection_error" {
			return echo.NewHTTPError(http.StatusServiceUnavailable, SERVICE_ERROR)
		} else {
			return echo.NewHTTPError(http.StatusNotFound, NOT_FOUND_ERROR)
		}
	}

	// Return success response
	return c.JSON(http.StatusOK, data)
}

func AddyMessage(c echo.Context) error {
	// Convert id param to int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NO_INT_ERROR)
	}
	// Retrieve message by id
	message, err := db.GetMessageById(id)
	if err != nil {
		if err.Error() == "connection_error" {
			return echo.NewHTTPError(http.StatusServiceUnavailable, SERVICE_ERROR)
		} else {
			return echo.NewHTTPError(http.StatusNotFound, NOT_FOUND_ERROR)
		}
	}

	message.Ys = message.Ys + 1

	// Update message in database
	data, err := db.UpdateMessage(message)
	if err != nil {
		if err.Error() == "connection_error" {
			return echo.NewHTTPError(http.StatusServiceUnavailable, SERVICE_ERROR)
		} else {
			return echo.NewHTTPError(http.StatusNotFound, NOT_FOUND_ERROR)
		}
	}

	// Return success response
	return c.JSON(http.StatusOK, data)
}

func Readiness(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
