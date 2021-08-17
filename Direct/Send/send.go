package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// send to queue
func main() {

	// connect to MQ
	ch := connectMQ()

	// declare a queue
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	body := fmt.Sprintf("Now is %v, the message is %v", time.Now(), "Hello Queue")

	// push a message
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if err != nil {
		log.Fatalf("%s: %s", "Failed to publish a message", err)
	}
}

// connectMQ 連線到 MQ
func connectMQ() *amqp.Channel {
	// MQ Connect
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	// defer conn.Close()

	// open a MQ channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}

	// defer ch.Close()

	return ch
}
