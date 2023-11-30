package rabbitmq

import (
    "context"
    "encoding/json"
    "log"
    "math/rand"
    "time"

    amqp "github.com/rabbitmq/amqp091-go"
)

type Orange struct {
    Size int `json:"size"`
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

func publishMessage(ch *amqp.Channel, q amqp.Queue, body []byte) error {
    return ch.PublishWithContext(
        context.Background(), // context
        "",                   // exchange
        q.Name,               // routing key
        false,                // mandatory
        false,                // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        })
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

    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    log.Printf("generating oranges...")
    for {
        orange := Orange{Size: r.Intn(30)}
        body, err := json.Marshal(orange)
        failOnError(err, "Failed to encode body")

        err = publishMessage(ch, q, body)
        failOnError(err, "Failed to publish a message")

        time.Sleep(1 * time.Second)
    }
}