package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/jorgeAM/basicGraphql/generated"
	"github.com/jorgeAM/basicGraphql/models"
)

func (r *mutationResolver) SignUp(ctx context.Context, input models.SignUpInput) (*models.User, error) {
	_, err := r.UserResolver.FindByEmail(input.Email)

	if err == nil {
		return nil, errors.New("This email already taken")
	}

	if input.Password != input.ConfirmPassword {
		return nil, errors.New("Passwords does not matches")
	}

	u := &models.User{
		Name:     *input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	err = u.HashPassword()

	if err != nil {
		return nil, err
	}

	return r.UserResolver.Create(u)
}

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.User, error) {
	u, err := r.UserResolver.FindByEmail(input.Email)

	if err != nil {
		return nil, errors.New("User does not exist")
	}

	err = u.CheckPassword(input.Password)

	if err != nil {
		return nil, errors.New("Password is incorrect")
	}

	return u, nil
}

func (r *queryResolver) Me(ctx context.Context, input models.MeInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
