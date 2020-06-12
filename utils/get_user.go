package utils

import (
	"context"
	"errors"

	"github.com/jorgeAM/basicGraphql/models"
	userresolver "github.com/jorgeAM/basicGraphql/repositories/user"
)

// GetUserFromContext return an instance of models.User struct
func GetUserFromContext(ctx context.Context, userRepository userresolver.Handler) (*models.User, error) {
	id, ok := ctx.Value(UserID).(string)

	if !ok || id == "" {
		return nil, errors.New("There is no id in context")
	}

	return userRepository.FindByID(id)
}

// GetUserIDFromContext return userID from context
func GetUserIDFromContext(ctx context.Context) (*string, error) {
	id, ok := ctx.Value(UserID).(string)

	if !ok || id == "" {
		return nil, errors.New("There is no id in context")
	}

	return &id, nil
}
