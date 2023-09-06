package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/liviudnicoara/go-graphql-poc/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, cmd model.CreateTodo) (*model.CmdResponse, error) {
	err := r.TodoRepository.Add(model.Todo{
		Text:   cmd.Text,
		UserID: cmd.UserID,
	})

	if err != nil {
		return &model.CmdResponse{
			IsSucessfull: false,
		}, nil
	}

	return &model.CmdResponse{
		IsSucessfull: true,
	}, nil
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, cmd model.UpdateTodo) (*model.CmdResponse, error) {
	panic(fmt.Errorf("not implemented: UpdateTodo - updateTodo"))
}

// CompleteTodo is the resolver for the completeTodo field.
func (r *mutationResolver) CompleteTodo(ctx context.Context, id string) (*model.CmdResponse, error) {
	panic(fmt.Errorf("not implemented: CompleteTodo - completeTodo"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := r.Resolver.TodoRepository.Get()

	if err != nil {
		return nil, err
	}

	// Convert to a slice of pointers
	results := make([]*model.Todo, len(todos))
	for i := range todos {
		results[i] = &todos[i]
	}
	return results, err
}

// TodosByUserID is the resolver for the todosByUserID field.
func (r *queryResolver) TodosByUserID(ctx context.Context, userID string) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: TodosByUserID - todosByUserID"))
}

// TodoByID is the resolver for the todoByID field.
func (r *queryResolver) TodoByID(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: TodoByID - todoByID"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// UsersByID is the resolver for the usersByID field.
func (r *queryResolver) UsersByID(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UsersByID - usersByID"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }