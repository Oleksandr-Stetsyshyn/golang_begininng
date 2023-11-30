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

func createConnection() (*amqp.Connection, error) {
    return amqp.Dial("amqp://guest:guest@localhost:5672/")
}

func createChannel(conn *amqp.Connection) (*amqp.Channel, error) {
    return conn.Channel()
}

func declareQueue(ch *amqp.Channel) (amqp.Queue, error) {
    return ch.QueueDeclare(
        "oranges", // name
        false,     // durable
        false,     // delete when unused
        false,     // exclusive
        false,     // no-wait
        nil,       // arguments
    )
}

func consumeMessages(ch *amqp.Channel, q amqp.Queue) (<-chan amqp.Delivery, error) {
    return ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
}

func Main() {
    conn, err := createConnection()
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := createChannel(conn)
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    q, err := declareQueue(ch)
    failOnError(err, "Failed to declare a queue")

    msgs, err := consumeMessages(ch, q)
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

	for range time.Tick(10 * time.Second) {
		log.Printf("Small: %d\n", basket.Small)
		log.Printf("Medium: %d\n", basket.Medium)
		log.Printf("Large: %d\n", basket.Large)
	}
}