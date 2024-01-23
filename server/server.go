package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"handler"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// https://echo.labstack.com/docs/middleware/cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Routes
	e.GET("/api/messages", handler.GetAllMessages)
	e.GET("/api/messages/:id", handler.GetMessage)
	e.POST("/api/messages", handler.CreateMessage)
	e.PUT("/api/messages/:id/like", handler.LikeMessage)
	e.PUT("/api/messages/:id/unlike", handler.UnlikeMessage)
	e.PUT("/api/messages/:id/y", handler.AddyMessage)

	// Service start at port :1323
	e.Logger.Fatal(e.Start(":1323"))
}
