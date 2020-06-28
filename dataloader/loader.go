package dataloader

import (
	"context"
	"time"

	"github.com/jorgeAM/basicGraphql/models"
	"github.com/jorgeAM/basicGraphql/repository"
	"github.com/jorgeAM/basicGraphql/utils"
)

const (
	maxBatch = 100
	wait     = 1 * time.Millisecond
)

// Loader handles all dataloaders
type Loader struct {
	UserLoader      *UserLoader
	TodoSliceLoader *TodoSliceLoader
}

func newUserLoaderConfig(repository *repository.Layer) UserLoaderConfig {
	return UserLoaderConfig{
		MaxBatch: maxBatch,
		Wait:     wait,
		Fetch: func(ids []int) ([]*models.User, []error) {
			users, err := repository.UserRepository.FindByIds(ids)

			if err != nil {
				return nil, []error{err}
			}

			userByID := make(map[int]*models.User)

			for _, user := range users {
				userByID[user.ID] = user
			}

			result := []*models.User{}

			for _, id := range ids {
				result = append(result, userByID[id])
			}

			return result, nil
		},
	}
}

func newTodoSliceLoaderConfig(repository *repository.Layer) TodoSliceLoaderConfig {
	return TodoSliceLoaderConfig{
		MaxBatch: maxBatch,
		Wait:     wait,
		Fetch: func(ids []int) ([][]*models.Todo, []error) {
			todos, err := repository.TodoRepository.FindByUserIds(ids)

			if err != nil {
				return nil, []error{err}
			}

			todosByID := make(map[int][]*models.Todo)

			for _, todo := range todos {
				todosByID[todo.UserID] = append(todosByID[todo.UserID], todo)
			}

			var result [][]*models.Todo

			for _, id := range ids {
				result = append(result, todosByID[id])
			}

			return result, nil
		},
	}
}

// NewLoader return a new instance of Loader struct
func NewLoader(repository *repository.Layer) *Loader {
	userLoaderConfig := newUserLoaderConfig(repository)
	todoSliceLoaderConfig := newTodoSliceLoaderConfig(repository)

	return &Loader{
		UserLoader:      NewUserLoader(userLoaderConfig),
		TodoSliceLoader: NewTodoSliceLoader(todoSliceLoaderConfig),
	}
}

// GetLoader returns loader from context
func GetLoader(ctx context.Context) *Loader {
	return ctx.Value(utils.Loaders).(*Loader)
}
