package src

import "fmt"

type Subscriber struct {
	publisher    *Publisher
	Subscription int
}

func (s *Subscriber) Init(p *Publisher) {
	s.publisher = p
}

func (s *Subscriber) Shutdown() {
	s.publisher.Unsubscribe(s.Subscription)
}

func (s *Subscriber) Speak() {
	fmt.Println("Hello")
}
