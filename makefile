graphql-init:
	go run github.com/99designs/gqlgen init
graphql-generate:
	go run -mod=mod github.com/99designs/gqlgen generate