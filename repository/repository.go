package repository

import (
	"github.com/jorgeAM/basicGraphql/db"

	todorepository "github.com/jorgeAM/basicGraphql/repository/todo"
	userrepository "github.com/jorgeAM/basicGraphql/repository/user"
)

// Layer handles all repositories
type Layer struct {
	UserRepository userrepository.Handler
	TodoRepository todorepository.Handler
}

// NewRepositoryLayer creates new instance of Layer
func NewRepositoryLayer(engine db.TYPE, dbhandler db.Handler) (*Layer, error) {
	userRepo, err := userrepository.NewUserRepository(engine, dbhandler)

	if err != nil {
		return nil, err
	}

	todoRepo, err := todorepository.NewTodoRepository(engine, dbhandler)

	if err != nil {
		return nil, err
	}

	return &Layer{
		UserRepository: userRepo,
		TodoRepository: todoRepo,
	}, nil
}
