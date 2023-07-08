package main

import (
	"background_rabbitmq/config"
	"background_rabbitmq/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.DatabaseConnection()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/message", handlers.MessageStreamer)
	e.Logger.Fatal(e.Start(":1323"))

}
