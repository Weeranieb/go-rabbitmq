# Simple Producer/Consumer Example

This example demonstrates the basic RabbitMQ messaging pattern with a producer that sends messages and a consumer that receives them.

## Prerequisites

1. RabbitMQ running (use `docker-compose up -d` from project root)
2. Go 1.16 or later

## Setup

1. Initialize Go module:

   ```bash
   go mod init github.com/Weeranieb/go-rabbitmq
   ```

2. Install RabbitMQ Go client:
   ```bash
   go get github.com/rabbitmq/amqp091-go
   ```

## Running the Example

### Terminal 1 - Start the Consumer

```bash
go run main.go consumer
```

### Terminal 2 - Start the Producer

```bash
go run main.go producer
```

## What Happens

1. The producer connects to RabbitMQ and declares a queue named "hello"
2. The producer sends 5 messages with 1-second intervals
3. The consumer connects to the same queue and receives all messages
4. Messages are automatically acknowledged (auto-ack=true)

## Expected Output

**Producer:**

```
Starting Producer...
 [x] Sent Hello World! Message #1
 [x] Sent Hello World! Message #2
 [x] Sent Hello World! Message #3
 [x] Sent Hello World! Message #4
 [x] Sent Hello World! Message #5
```

**Consumer:**

```
Starting Consumer...
 [*] Waiting for messages. To exit press CTRL+C
 [x] Received Hello World! Message #1
 [x] Received Hello World! Message #2
 [x] Received Hello World! Message #3
 [x] Received Hello World! Message #4
 [x] Received Hello World! Message #5
```

## Key Concepts

- **Queue**: A buffer that stores messages
- **Producer**: Sends messages to a queue
- **Consumer**: Receives messages from a queue
- **Channel**: A virtual connection inside a connection
- **Auto-ack**: Messages are automatically acknowledged when received
