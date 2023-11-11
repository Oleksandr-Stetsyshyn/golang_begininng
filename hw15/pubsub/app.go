package fileChangedEvent

import (
	"fmt"
)

func Main() {
	publisher := &Publisher{}

	Bob := &Subscriber{Events: make(chan string)}
	publisher.Subscribe(Bob)

	go func() {
		for event := range Bob.Events {
			fmt.Println("Bob notified about:", event)
		}
	}()

	Earl := &Subscriber{Events: make(chan string)}
	publisher.Subscribe(Earl)

	go func() {
		for event := range Earl.Events {
			fmt.Println("Earl notified about:", event)
		}
	}()

	Jon := &Subscriber{Events: make(chan string)}
	publisher.Subscribe(Jon)

	go func() {
		for event := range Jon.Events {
			fmt.Println("Jon notified about", event)
		}
	}()

	FolderWatcher("hw15/pubsub/files", publisher)
}
