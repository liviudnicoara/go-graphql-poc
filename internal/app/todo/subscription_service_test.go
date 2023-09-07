package todo

import (
	"testing"

	"github.com/google/uuid"
	"github.com/liviudnicoara/go-graphql-poc/shared"
)

func TestSubscribeUnsubscribe(t *testing.T) {
	service := NewSubscriptionService()
	userID := uuid.New()

	// Subscribe
	subscriber := service.Subscribe(userID)

	// Ensure subscriber exists
	_, exists := service.subscribers[userID]
	if !exists {
		t.Errorf("Expected subscriber with ID %s to exist", userID)
	}

	// Unsubscribe
	service.Unsubscribe(userID)

	// Ensure subscriber is removed
	_, exists = service.subscribers[userID]
	if exists {
		t.Errorf("Expected subscriber with ID %s to be removed", userID)
	}

	// Ensure the subscriber channel is closed
	select {
	case _, open := <-subscriber:
		if open {
			t.Errorf("Subscriber channel should be closed after unsubscribe")
		}
	default:
	}
}

func TestPublish(t *testing.T) {
	service := NewSubscriptionService()
	userID := uuid.New()
	todo := shared.Todo{ID: uuid.New().String(), Text: "Test Todo", UserID: userID.String()}

	// Subscribe
	subscriber := service.Subscribe(userID)

	// Publish an event
	service.Publish(todo)

	// Receive the event from the subscriber
	receivedTodo := <-subscriber

	if receivedTodo != todo {
		t.Errorf("Expected to receive Todo %+v; got %+v", todo, receivedTodo)
	}
}

func TestStop(t *testing.T) {
	service := NewSubscriptionService()
	userID := uuid.New()

	// Subscribe
	service.Subscribe(userID)

	// Stop the service
	service.Stop()

	// Ensure all subscribers are removed
	_, exists := service.subscribers[userID]
	if exists {
		t.Errorf("Expected all subscribers to be removed after stopping the service")
	}
}
