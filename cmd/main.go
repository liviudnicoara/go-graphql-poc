package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/liviudnicoara/go-graphql-poc/internal/app/todo"
	"github.com/liviudnicoara/go-graphql-poc/internal/app/user"
	"github.com/liviudnicoara/go-graphql-poc/internal/transport/graphql"
)

const defaultPort = "8080"

func main() {

	subscriptionService := todo.NewSubscriptionService()
	defer subscriptionService.Stop()
	todoRepo := todo.NewToDoRepository()

	// Start the background task in a goroutine
	go func() {
		// Define the interval at which to check for overdue todos
		interval := 10 * time.Second // Check daily, adjust as needed

		// Create a ticker to trigger the function at the specified interval
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		// Run the function immediately when the application starts
		checkOverdueTodos(subscriptionService, todoRepo)

		// Enter a loop to repeatedly run the function
		for {
			select {
			case <-ticker.C:
				// Time to run the function again
				checkOverdueTodos(subscriptionService, todoRepo)
			}
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := graphql.Resolver{
		UserRepo:        user.NewUserRepository(),
		TodoRepo:        todoRepo,
		SubscriptionSvc: subscriptionService,
	}

	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: &resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func checkOverdueTodos(svc todo.SubscriptionService, todorepo todo.TodoRepository) {

	todos, err := todorepo.Get()
	log.Println(todos)

	if err != nil || len(todos) == 0 {
		return
	}

	// now := time.Now()

	for _, todo := range todos {
		// if todo.DueAt.Before(now) {
		svc.Publish(todo)
		// }
	}
}
