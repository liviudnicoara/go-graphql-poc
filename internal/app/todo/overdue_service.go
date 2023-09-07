package todo

import (
	"log"
	"time"
)

type OverDueService interface {
	Start()
	Stop()
}

type overDueService struct {
	subSvc   SubscriptionService
	todoRepo TodoRepository
	stopCh   chan struct{}
}

// NewOverDueService creates a new instance of the OverDueService.
func NewOverDueService(subSvc SubscriptionService, todoRepo TodoRepository) OverDueService {
	return &overDueService{
		subSvc:   subSvc,
		todoRepo: todoRepo,
		stopCh:   make(chan struct{}),
	}
}

func (svc *overDueService) Start() {
	// Start the background task in a goroutine
	go func(stopCh chan struct{}) {
		// Define the interval at which to check for overdue todos
		interval := 10 * time.Second // Check daily, adjust as needed

		// Create a ticker to trigger the function at the specified interval
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		// Run the function immediately when the application starts
		svc.checkOverdueTodos()

		// Enter a loop to repeatedly run the function
		for {
			select {
			case <-ticker.C:
				// Time to run the function again
				svc.checkOverdueTodos()
			case <-stopCh:
				return
			}
		}
	}(svc.stopCh)
}

func (svc *overDueService) Stop() {
	close(svc.stopCh)
}

func (svc *overDueService) checkOverdueTodos() {

	todos, err := svc.todoRepo.Get()
	log.Println(todos)

	if err != nil || len(todos) == 0 {
		return
	}

	now := time.Now()

	for _, todo := range todos {
		if todo.DueAt.Before(now) {
			svc.subSvc.Publish(todo)
		}
	}
}
