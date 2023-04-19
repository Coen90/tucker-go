package main

import "context"

type Publisher struct {
	ctx context.Context
	subscribeCh chan chan<- string // 데이터를 넣기만 할 수 있는 채널
	publisherCh chan string
	subscribers []chan<- string
}

func NewPublisher(ctx context.Context) *Publisher {
	return &Publisher{
		ctx: ctx,
		subscribeCh: make(chan chan<- string),
		publisherCh: make(chan string),
		subscribers: make([]chan<- string, 0),
	}
}

func (p *Publisher) Subscribe(sub chan<- string) {
	p.subscribeCh <- sub
}

func (p *Publisher) Publish(msg string) {
	p.publisherCh <- msg
}

func (p *Publisher) Update() {
	for {
		select {
		case sub := <-p.subscribeCh:
			p.subscribers = append(p.subscribers, sub)
		case msg := <-p.publisherCh:
			for _, subscriber := range p.subscribers {
				subscriber <- msg
			}
		case <-p.ctx.Done():
			wg.Done()
			return
		}
	}
}