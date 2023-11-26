
package rabbitmqListener

import (
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Orange struct {
	Size int `json:"size"`
}

type Basket struct {
	Small  int
	Medium int
	Large  int
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"oranges", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	basket := Basket{}

	go func() {
		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		for d := range msgs {
			orange := Orange{}
			err := json.Unmarshal(d.Body, &orange)
			failOnError(err, "Failed to decode orange")

			switch {
			case orange.Size < 10:
				basket.Small++
			case orange.Size < 20:
				basket.Medium++
			default:
				basket.Large++
			}
		}
	}()

	for range time.Tick(1 * time.Minute) {
		log.Printf("Small: %d\n", basket.Small)
		log.Printf("Medium: %d\n", basket.Medium)
		log.Printf("Large: %d\n", basket.Large)
	}
}
