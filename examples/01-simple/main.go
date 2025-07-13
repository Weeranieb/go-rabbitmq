package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	queueName = "hello"
	brokerURL = "amqp://user:password@localhost:5672/"
)

// Producer sends messages to RabbitMQ
func producer() {
	// Connect to RabbitMQ
	conn, err := amqp.Dial(brokerURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare a queue
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Send messages
	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Hello World! Message #%d", i)

		err = ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			})

		if err != nil {
			log.Fatalf("Failed to publish a message: %v", err)
		}

		log.Printf(" [x] Sent %s", message)
		time.Sleep(1 * time.Second)
	}
}

// Consumer receives messages from RabbitMQ
func consumer() {
	// Connect to RabbitMQ
	conn, err := amqp.Dial(brokerURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare a queue
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Consume messages
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	for {
		select {
		case msg := <-msgs:
			log.Printf(" [x] Received %s", msg.Body)
		case <-sigChan:
			log.Println("Shutting down consumer...")
			return
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [producer|consumer]")
		fmt.Println("  producer - Send messages to RabbitMQ")
		fmt.Println("  consumer - Receive messages from RabbitMQ")
		os.Exit(1)
	}

	mode := os.Args[1]

	switch mode {
	case "producer":
		log.Println("Starting Producer...")
		producer()
	case "consumer":
		log.Println("Starting Consumer...")
		consumer()
	default:
		fmt.Printf("Unknown mode: %s\n", mode)
		fmt.Println("Available modes: producer, consumer")
		os.Exit(1)
	}
}
