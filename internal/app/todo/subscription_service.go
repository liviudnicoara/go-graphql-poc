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
	Stop()
}

type subscriptionService struct {
	subscribers map[uuid.UUID]chan shared.Todo
	mutex       sync.RWMutex
	stopCh      chan struct{} // Channel to signal goroutines to stop
	wg          sync.WaitGroup
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

	ps.wg.Add(1)

	go func(ch chan shared.Todo) {
		defer ps.wg.Done()

		select {
		case ch <- event:
		case <-ps.stopCh:
			return
		}
	}(ps.subscribers[uuid.MustParse(event.UserID)])

}

func (ps *subscriptionService) Stop() {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	close(ps.stopCh) // Signal goroutines to stop

	// Wait for all goroutines to complete
	ps.wg.Wait()

	// Close remaining event channels
	for id, ch := range ps.subscribers {
		close(ch)
		delete(ps.subscribers, id)
	}
}
