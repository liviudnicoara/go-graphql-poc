package main

import (
	"log"
	"net/http"
	"os"

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

	overDueService := todo.NewOverDueService(subscriptionService, todoRepo)
	overDueService.Start()
	defer overDueService.Stop()

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
