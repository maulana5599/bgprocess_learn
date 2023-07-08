package producers

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
)

func MessagePublisher(messageStream string, c echo.Context) error {
	connectRabbitMQ, channelRabbitMQ := sender()

	defer connectRabbitMQ.Close()

	defer channelRabbitMQ.Close()

	_, errCh := channelRabbitMQ.QueueDeclare(
		"MessageQueue", // queue name
		true,           // durable
		false,          // auto delete
		false,          // exclusive
		false,          // no wait
		nil,            // arguments
	)
	if errCh != nil {
		panic(errCh)
	}

	// Create a message to publish.
	message := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(messageStream),
	}

	// Attempt to publish a message to the queue.
	if err := channelRabbitMQ.Publish(
		"",             // exchange
		"MessageQueue", // queue name
		false,          // mandatory
		false,          // immediate
		message,        // message to nilpublish
	); err != nil {
		log.Println(err.Error())
		return err
	}

	return c.JSON(200, echo.Map{
		"message": messageStream,
	})
}

func sender() (*amqp.Connection, *amqp.Channel) {

	amqpServerURL := "amqp://guest:guest@localhost:5672/"

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return connectRabbitMQ, channelRabbitMQ
}
