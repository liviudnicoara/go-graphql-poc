package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/liviudnicoara/go-graphql-poc/graph"
	"github.com/liviudnicoara/go-graphql-poc/graph/model"
	"github.com/liviudnicoara/go-graphql-poc/repositories"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := graph.Resolver{
		UserRepository: repositories.NewBaseRepository[model.User](),
		TodoRepository: repositories.NewBaseRepository[model.Todo](),
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
