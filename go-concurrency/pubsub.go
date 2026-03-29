package main

import (
	"fmt"
	"sync"
)

type Message struct {
	Type string
	Data string
}

type Subscribers map[string][]chan Message

type Broker struct {
	Subscribers Subscribers
	rwMu        sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		Subscribers: make(map[string][]chan Message),
	}
}

func (b *Broker) Subscribe(topic string) <-chan Message {
	channel := make(chan Message, 10)

	b.rwMu.Lock()
	b.Subscribers[topic] = append(b.Subscribers[topic], channel)
	b.rwMu.Unlock()

	return channel
}

func (b *Broker) Publish(topic string, msg Message) {
	b.rwMu.RLock()
	defer b.rwMu.RUnlock()

	for _, ch := range b.Subscribers[topic] {
		ch <- msg
	}
}

func (b *Broker) Unsubscribe(topic string, sub <-chan Message) {
	b.rwMu.Lock()
	defer b.rwMu.Unlock()

	subs, ok := b.Subscribers[topic]
	if !ok {
		return
	}

	for idx, subscriber := range subs {
		if subscriber == sub {
			b.Subscribers[topic] = append(subs[:idx], subs[idx+1:]...)
			close(subscriber)
		}
	}

	if len(b.Subscribers) == 0 {
		delete(b.Subscribers, topic)
	}
}

func StartPubSub() {
	var wg sync.WaitGroup

	broker := NewBroker()

	sub1 := broker.Subscribe("cricket")
	sub2 := broker.Subscribe("football")

	wg.Add(2)

	go func() {
		wg.Done()
		for msg := range sub1 {
			fmt.Println(msg.Data)
		}
	}()

	go func() {
		wg.Done()
		for msg := range sub2 {
			fmt.Println(msg.Data)
		}
	}()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			broker.Publish("cricket", Message{
				Type: "text",
				Data: "virat hit 6",
			})
		} else {
			broker.Publish("football", Message{
				Type: "text",
				Data: "ronaldo scored a goal",
			})
		}
	}

	broker.Unsubscribe("cricket", sub1)
	broker.Unsubscribe("football", sub2)

	wg.Wait()
}
