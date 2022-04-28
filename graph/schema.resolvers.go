package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/deb151292/gqlgen-todos/graph/generated"
	"github.com/deb151292/gqlgen-todos/graph/logics"
	"github.com/deb151292/gqlgen-todos/graph/model"
)

func (r *queryResolver) CreateQR(ctx context.Context, input model.InputStr) (*model.SuccessMsg, error) {
	// var m *model.SuccessMsg
	// err := errors.New("abcd")

	m, err := logics.QrGenerator(input)
	return m, err
}

func (r *queryResolver) FileMove(ctx context.Context, input *model.Paths) (*model.Response, error) {
	massage, err := logics.Moving(input)
	return massage, err
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
