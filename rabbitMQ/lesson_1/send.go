package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

// We also need a helper function to check the return value for each amqp call:
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	//connect to rabbitMQ server
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Failed to connect to rabbitMQ")
	defer conn.Close()

	//The connection abstracts the socket connection, and takes care of protocol version negotiation and authentication and so on for us.
	//Next we create a channel, which is where most of the API for getting things done resides:
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//To send, we must declare a queue for us to send to; then we can publish a message to the queue:
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
