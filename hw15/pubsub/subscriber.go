package fileChangedEvent

import "fmt"

type Subscriber struct {
	name   string
	Events chan string
}
func NewSubscriber(name string) *Subscriber {
    s := &Subscriber{
        name:   name,
        Events: make(chan string),
    }

    go s.Listen()
    return s
}

func (s *Subscriber) Listen() {
	for event := range s.Events {
		fmt.Println(s.name, "notified about:", event)
	}
}
