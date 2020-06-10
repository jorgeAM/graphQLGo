package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jorgeAM/basicGraphql/generated"
	"github.com/jorgeAM/basicGraphql/models"
)

func (r *mutationResolver) SignUp(ctx context.Context, input models.SignUpInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context, input models.MeInput) (*models.User, error) {
	return &models.User{
		ID:       input.ID,
		Email:    "jorge.alfmur@gmail.com",
		Name:     "Jorguito",
		Password: "123456",
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
