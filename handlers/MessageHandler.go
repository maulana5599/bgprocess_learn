package handlers

import (
	"background_rabbitmq/producers"

	"github.com/labstack/echo/v4"
)

func MessageStreamer(ctx echo.Context) error {
	message := ctx.QueryParam("message")
	return producers.MessagePublisher(message, ctx)
}
