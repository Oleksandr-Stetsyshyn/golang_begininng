package fileChangedEvent

import "fmt"

func Main() {
	broker := &Broker{EventChannel: make(chan string)}
	publisher := &Publisher{Broker: broker, Path: "hw15/pubsub/files"}

	Bob := NewSubscriber("Bob")
	broker.Subscribe(Bob)

	Alice := NewSubscriber("Alice")
	broker.Subscribe(Alice)

	go publisher.FolderWatcher()
	fmt.Scanln()
}
