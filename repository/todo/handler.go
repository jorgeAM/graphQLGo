package todorepository

import (
	"fmt"

	"github.com/jorgeAM/basicGraphql/db"
	"github.com/jorgeAM/basicGraphql/models"
)

// Handler interface has all method to userRepository
type Handler interface {
	Create(todo *models.Todo) (*models.Todo, error)
	FindByUserIds(userIds []int) ([]*models.Todo, error)
}

// NewTodoRepository initialize userRepository
func NewTodoRepository(engine db.TYPE, dbhandler db.Handler) (Handler, error) {
	switch engine {
	case db.POSTGRES:
		psqlHandler := dbhandler.(*db.PostgresHandler)
		return &PSQLRepository{psqlHandler}, nil
	default:
		return nil, fmt.Errorf("%s engine is not supported yet", engine)
	}
}
