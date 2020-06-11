package userrepository

import (
	"fmt"

	"github.com/jorgeAM/basicGraphql/db"
	"github.com/jorgeAM/basicGraphql/models"
)

// Handler interface has all method to userRepository
type Handler interface {
	Create(user *models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

// NewUserRepository initialize userRepository
func NewUserRepository(engine db.TYPE, dbhandler db.Handler) (Handler, error) {
	switch engine {
	case db.POSTGRES:
		psqlHandler := dbhandler.(*db.PostgresHandler)
		return &PSQLRepository{psqlHandler}, nil
	default:
		return nil, fmt.Errorf("%s engine is not supported yet", engine)
	}
}
