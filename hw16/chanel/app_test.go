package rabbitmq

import (
    "testing"
)

func TestCreateConnection(t *testing.T) {
    _, err := createConnection()
    if err != nil {
        t.Errorf("Expected no error for createConnection, got %v", err)
    }
}

func TestCreateChannel(t *testing.T) {
    conn, _ := createConnection()
    _, err := createChannel(conn)
    if err != nil {
        t.Errorf("Expected no error for createChannel, got %v", err)
    }
}

func TestDeclareQueue(t *testing.T) {
    conn, _ := createConnection()
    ch, _ := createChannel(conn)
    _, err := declareQueue(ch)
    if err != nil {
        t.Errorf("Expected no error for declareQueue, got %v", err)
    }
}

func TestPublishMessage(t *testing.T) {
    conn, _ := createConnection()
    ch, _ := createChannel(conn)
    q, _ := declareQueue(ch)
    body := []byte("test message")
    err := publishMessage(ch, q, body)
    if err != nil {
        t.Errorf("Expected no error for publishMessage, got %v", err)
    }
}