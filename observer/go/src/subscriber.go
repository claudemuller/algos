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
	fmt.Println("Subscriber: Unsubscribing...")
	s.publisher.Unsubscribe(s.Subscription)
}

func (s *Subscriber) Speak(d string) {
	fmt.Printf("Subscriber: Publisher says \"%s\"\n", d)
}
