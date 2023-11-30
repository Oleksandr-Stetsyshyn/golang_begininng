package rabbitmqListener

import (
    "testing"

    amqp "github.com/rabbitmq/amqp091-go"
)

// Test for createConnection function
func TestCreateConnection(t *testing.T) {
    _, err := createConnection()
    if err != nil {
        t.Errorf("Failed to create connection: %s", err)
    }
}

// Test for createChannel function
func TestCreateChannel(t *testing.T) {
    conn, _ := createConnection()
    _, err := createChannel(conn)
    if err != nil {
        t.Errorf("Failed to create channel: %s", err)
    }
}

// Test for declareQueue function
func TestDeclareQueue(t *testing.T) {
    conn, _ := createConnection()
    ch, _ := createChannel(conn)
    _, err := declareQueue(ch)
    if err != nil {
        t.Errorf("Failed to declare queue: %s", err)
    }
}

// Negative test for createConnection function
func TestCreateConnectionNegative(t *testing.T) {
    // Trying to connect to a non-existing server
    conn, err := amqp.Dial("amqp://guest:guest@nonexistingserv:5672/")
    if err == nil {
        t.Errorf("Expected error, got nil")
    }
    if conn != nil {
        t.Errorf("Expected nil, got %v", conn)
    }
}