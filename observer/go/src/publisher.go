package src

import "time"

type Publisher struct {
	subscriptions []func()
}

func (p *Publisher) Start() {
	time.Sleep(time.Microsecond * 5)
	p.end()
}

func (p *Publisher) end() {
	p.notify()
}

func (p *Publisher) notify() {
	for _, s := range p.subscriptions {
		s()
	}
}

func (p *Publisher) Subscribe(f func()) int {
	p.subscriptions = append(p.subscriptions, f)

	return len(p.subscriptions) - 1
}

func (p *Publisher) Unsubscribe(i int) {
	p.subscriptions = append(p.subscriptions[:i], p.subscriptions[i+1:len(p.subscriptions)]...)
}
