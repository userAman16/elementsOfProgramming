package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Message struct {
	Topic   string
	Payload any
}

type Subscriber chan Message

type PubSub struct {
	mu          sync.RWMutex
	subscribers map[string][]Subscriber
	ctx         context.Context
	cancel      context.CancelFunc
}

func NewPubSub() *PubSub {
	ctx, cancel := context.WithCancel(context.Background())
	return &PubSub{
		subscribers: make(map[string][]Subscriber),
		ctx:         ctx,
		cancel:      cancel,
	}
}

func (ps *PubSub) Subscribe(topic string, buffer int) Subscriber {
	sub := make(Subscriber, buffer)
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.subscribers[topic] = append(ps.subscribers[topic], sub)
	return sub
}

func (ps *PubSub) UnSubscribe(topic string, sub Subscriber) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	subs := ps.subscribers[topic]
	for i, s := range subs {
		if sub == s {
			ps.subscribers[topic] = append(ps.subscribers[topic][:i], ps.subscribers[topic][i+1:]...)
			close(sub)
			break
		}
	}

}

func (ps *PubSub) Publish(topic string, payload any) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	subs := ps.subscribers[topic]

	for _, s := range subs {
		select {
		case s <- Message{Topic: topic, Payload: payload}:
		}

	}
}

func (ps *PubSub) Shutdown() {
	ps.cancel()
	ps.mu.Lock()
	defer ps.mu.Unlock()
	for topic, subs := range ps.subscribers {
		for _, sub := range subs {
			close(sub)
		}
		ps.subscribers[topic] = nil
	}
}

func PubSubMain() {
	pubSub := NewPubSub()
	defer pubSub.Shutdown()

	sub := pubSub.Subscribe("news", 10)

	go func() {
		for msg := range sub {
			fmt.Println("Received ", msg.Topic, msg.Payload)
		}
	}()

	pubSub.Publish("news", "Some Good news")
	time.Sleep(time.Second)

}
