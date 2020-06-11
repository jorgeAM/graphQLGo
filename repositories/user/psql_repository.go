package userrepository

import (
	"github.com/jorgeAM/basicGraphql/db"
	"github.com/jorgeAM/basicGraphql/models"
)

// PSQLRepository hits postgres database
type PSQLRepository struct {
	*db.PostgresHandler
}

// Create method create a new user in database
func (r *PSQLRepository) Create(user *models.User) (*models.User, error) {
	err := r.DB.Insert(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindByEmail method finds a user by its email
func (r *PSQLRepository) FindByEmail(email string) (*models.User, error) {
	user := new(models.User)
	err := r.DB.Model(user).Where("email = ?", email).First()

	if err != nil {
		return nil, err
	}

	return user, nil
}
