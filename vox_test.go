package vox_test

import (
	"testing"

	"github.com/gaurishhs/vox"
)

func TestPublisher(t *testing.T) {
	// Create a new publisher
	publisher := vox.NewPublisher()
	// Create a new subscriber
	subscriber := vox.NewSubscriber()
	// Subscribe to a topic
	subscriber.Subscribe("topic1")
	// Add the subscriber to the publisher
	publisher.AddSubscriber(subscriber)
	// Publish a message to the topic
	publisher.Publish(vox.NewMessage("topic1", map[string]interface{}{
		"message": "Hello World!",
	}))
	// Listen for messages on the messages channel
	subscriber.Listen(func(msg *vox.Message) {
		if msg.Topic != "topic1" {
			t.Errorf("Expected topic to be topic1, got %s", msg.Topic)
		}
		if msg.Payload["message"] != "Hello World!" {
			t.Errorf("Expected message to be Hello World!, got %s", msg.Payload["message"])
		}
		subscriber.Close()
	})
}
