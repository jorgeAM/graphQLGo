package todorepository

import (
	"github.com/jorgeAM/basicGraphql/db"
	"github.com/jorgeAM/basicGraphql/models"
)

// PSQLRepository hits postgres database
type PSQLRepository struct {
	*db.PostgresHandler
}

// Create method create a new todo in database
func (r *PSQLRepository) Create(todo *models.Todo) (*models.Todo, error) {
	err := r.DB.Insert(todo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
