package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/jorgeAM/basicGraphql/dataloader"
	"github.com/jorgeAM/basicGraphql/generated"
	"github.com/jorgeAM/basicGraphql/models"
	"github.com/jorgeAM/basicGraphql/utils"
)

func (r *mutationResolver) SignUp(ctx context.Context, input models.SignUpInput) (*models.Auth, error) {
	_, err := r.UserRepository.FindByEmail(input.Email)

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

	_, err = r.UserRepository.Create(u)

	if err != nil {
		return nil, err
	}

	return u.GenerateToken()
}

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.Auth, error) {
	u, err := r.UserRepository.FindByEmail(input.Email)

	if err != nil {
		return nil, errors.New("User does not exist")
	}

	err = u.CheckPassword(input.Password)

	if err != nil {
		return nil, errors.New("Password is incorrect")
	}

	return u.GenerateToken()
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input models.CreateTodoInput) (*models.Todo, error) {
	id, err := utils.GetUserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	casted, err := strconv.Atoi(*id)

	if err != nil {
		return nil, err
	}

	todo := &models.Todo{
		Title:       input.Title,
		Description: input.Description,
		UserID:      casted,
	}

	return r.TodoRepository.Create(todo)
}

func (r *queryResolver) Me(ctx context.Context) (*models.User, error) {
	return utils.GetUserFromContext(ctx, r.UserRepository)
}

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	return dataloader.GetLoader(ctx).UserLoader.Load(obj.UserID)
}

func (r *userResolver) Todos(ctx context.Context, obj *models.User) ([]*models.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
