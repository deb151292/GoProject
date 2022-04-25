package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/deb151292/gqlgen-todos/graph/generated"
	"github.com/deb151292/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	log.Println("abcd1")
	a := model.Todo{}
	return &a, errors.New("abcd")
}

func (r *mutationResolver) CreateQR(ctx context.Context, input string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	log.Println("abcd2")
	a := []*model.Todo{}
	return a, errors.New("abcd1")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
