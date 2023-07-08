package consumers

import (
	"background_rabbitmq/config"
	"background_rabbitmq/models"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func MessageConsumer() {
	// Define RabbitMQ server URL.
	amqpServerURL := "amqp://guest:guest@localhost:5672/"

	config.DatabaseConnection()

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	// Subscribing to QueueService2 for getting messages.
	messages, err := channelRabbitMQ.Consume(
		"MessageQueue", // queue name
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no local
		false,          // no wait
		nil,            // arguments
	)
	if err != nil {
		log.Println(err)
	}

	// Build a welcome message.
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)

	go func() {
		for message := range messages {
			fmt.Println(message.Body)
			// For example, show received message in a console.
			log.Printf(" > Received message: %s\n", message.Body)
			// for i := 0; i < 1; i++ {
			models.SaveMessage(string(message.Body))
			// }
		}
	}()

	<-forever
}
