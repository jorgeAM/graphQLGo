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
	UserLoader *UserLoader
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

// NewLoader return a new instance of Loader struct
func NewLoader(repository *repository.Layer) *Loader {
	userLoaderConfig := newUserLoaderConfig(repository)

	return &Loader{
		UserLoader: NewUserLoader(userLoaderConfig),
	}
}

// GetLoader returns loader from context
func GetLoader(ctx context.Context) *Loader {
	return ctx.Value(utils.Loaders).(*Loader)
}
