package vox

import (
	"sync"
	"time"
)

// Message represents a message published to a topic
type Message struct {
	// The topic the message is published to
	Topic string
	// The payload of the message
	Payload map[string]interface{}
}

// Create a new message
func NewMessage(topic string, payload map[string]interface{}) *Message {
	return &Message{
		Topic:   topic,
		Payload: payload,
	}
}

// Subscriber represents a subscriber to a topic
type Subscriber struct {
	// A unique identifier for the subscriber
	id int64
	// Messages channel
	messages chan *Message
	// Topics the subscriber is subscribed to
	topics map[string]bool
	// Lock
	lock sync.RWMutex
}

// Generate a unique ID for the subscriber
func generateUniqueID() int64 {
	return time.Now().UTC().UnixNano()
}

// Create a new subscriber
func NewSubscriber() *Subscriber {
	return &Subscriber{
		id:       generateUniqueID(),
		messages: make(chan *Message),
		topics:   make(map[string]bool),
	}
}

// Subscribe to a topic
func (s *Subscriber) Subscribe(topic string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.topics[topic] = true
}

// Unsubscribe from a topic
func (s *Subscriber) Unsubscribe(topic string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.topics, topic)
}

// Check if the subscriber is subscribed to a topic
func (s *Subscriber) IsSubscribed(topic string) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.topics[topic]
	return ok
}

// Get the messages channel
func (s *Subscriber) Messages() <-chan *Message {
	return s.messages
}

// Close the messages channel
func (s *Subscriber) Close() {
	close(s.messages)
}

// Get the subscriber ID
func (s *Subscriber) ID() int64 {
	return s.id
}

// Signal a message to the subscriber
func (s *Subscriber) Signal(message *Message) {
	s.messages <- message
}

// Listen for messages on the messages channel
func (s *Subscriber) Listen(callback func(*Message)) {
	for msg := range s.messages {
		callback(msg)
	}
}

// Publisher represents a publisher of messages to topics
type Publisher struct {
	// Lock
	mute sync.RWMutex
	// Subscribers
	subscribers map[int64]*Subscriber
}

// Create a new publisher
func NewPublisher() *Publisher {
	return &Publisher{
		subscribers: make(map[int64]*Subscriber),
	}
}

// Add a subscriber to the publisher
func (p *Publisher) AddSubscriber(subscriber *Subscriber) {
	p.mute.Lock()
	defer p.mute.Unlock()
	p.subscribers[subscriber.id] = subscriber
}

// Remove a subscriber from the publisher
func (p *Publisher) RemoveSubscriber(subscriber *Subscriber) {
	p.mute.Lock()
	defer p.mute.Unlock()
	delete(p.subscribers, subscriber.id)
}

// Publish a message to the topic
func (p *Publisher) Publish(message *Message) {
	p.mute.RLock()
	defer p.mute.RUnlock()
	for _, subscriber := range p.subscribers {
		if !subscriber.IsSubscribed(message.Topic) {
			continue
		}
		m := NewMessage(message.Topic, message.Payload)
		go (func(s *Subscriber) {
			s.Signal(m)
		})(subscriber)
	}
}
