package todo

import (
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/liviudnicoara/go-graphql-poc/shared"
)

type SubscriptionService interface {
	Subscribe(id uuid.UUID) chan shared.Todo
	Unsubscribe(id uuid.UUID)
	Publish(todo shared.Todo)
}

type subscriptionService struct {
	subscribers map[uuid.UUID]chan shared.Todo
	mutex       sync.RWMutex
}

// NewSubscriptionService creates a new instance of the SubscriptionService.
func NewSubscriptionService() *subscriptionService {
	return &subscriptionService{
		subscribers: make(map[uuid.UUID]chan shared.Todo),
	}
}

// Subscribe adds a new subscriber with a specified UUID.
func (ps *subscriptionService) Subscribe(id uuid.UUID) chan shared.Todo {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	eventChannel := make(chan shared.Todo)
	ps.subscribers[id] = eventChannel

	log.Println(id)
	log.Println(ps.subscribers)

	return eventChannel
}

// Unsubscribe removes a subscriber by its UUID.
func (ps *subscriptionService) Unsubscribe(id uuid.UUID) {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	if ch, ok := ps.subscribers[id]; ok {
		close(ch)
		delete(ps.subscribers, id)
	}
}

// Publish broadcasts an event to subscriber.
func (ps *subscriptionService) Publish(event shared.Todo) {
	ps.mutex.RLock()
	defer ps.mutex.RUnlock()

	log.Println(event)

	go func(ch chan shared.Todo) {
		log.Println("event: ", event)
		ch <- event
	}(ps.subscribers[uuid.MustParse(event.UserID)])

}
