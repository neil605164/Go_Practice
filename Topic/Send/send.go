package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func main() {
	// connect to MQ
	ch := connectMQ()

	// create a new exchange if exchange is not exist
	err := ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a exchange", err)
	}

	body := bodyFrom(os.Args)

	for i := 0; i <= 0; i++ {

		msg := fmt.Sprintf("%d,%v", i, body)

		err := ch.Publish(
			"logs_topic",          // exchange
			severityFrom(os.Args), // routing key
			false,                 // mandatory
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(msg),
			})

		if err != nil {
			log.Fatalf("%s: %s", "Failed to publish a message", err)
		}
		log.Printf(" [x] Num %d Sent %s", i, body)
	}
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "anonymous.info"
	} else {
		s = os.Args[1]
	}
	return s
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
