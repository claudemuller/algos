package main

import "github.com/claudemuller/algos/observer/go/src"

func main() {
	p := &src.Publisher{}
	s := &src.Subscriber{}

	// Store a pointer to the publisher on the subscriber
	s.Init(p)

	// Subscribe to publisher with subscriber's Speak func
	p.Subscribe(s.Speak)

	// Start publisher's time wasting
	p.Start()

	// Unsubscribe subscriber from publisher
	s.Shutdown()
}
