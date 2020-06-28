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

// FindByUserIds search todos per userId
func (r *PSQLRepository) FindByUserIds(userIds []int) ([]*models.Todo, error) {
	var todos []*models.Todo

	err := r.DB.Model(&todos).WhereIn("user_id in (?)", userIds).Select()

	if err != nil {
		return nil, err
	}

	return todos, nil
}
