package fileChangedEvent

type Broker struct {
	Subscribers  []*Subscriber
	EventChannel chan string
}

func (b *Broker) Subscribe(s *Subscriber) {
	b.Subscribers = append(b.Subscribers, s)
}

func (b *Broker) NotifyAll(event string) {
	for _, subscriber := range b.Subscribers {
		go func(s *Subscriber) {
			s.Events <- event
		}(subscriber)
	}
}
