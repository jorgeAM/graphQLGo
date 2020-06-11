package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" pg:"type:varchar(100)"`
	Email     string    `json:"email" pg:",unique,type:varchar(100)"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt" pg:"default:now()"`
	UpdatedAt time.Time `json:"updatedAt" pg:"default:now()"`
}

// HashPassword encrypt password
func (u *User) HashPassword() error {
	passBytes := []byte(u.Password)
	hashBytes, err := bcrypt.GenerateFromPassword(passBytes, bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashBytes)
	return nil
}

// CheckPassword verifies if password is correct
func (u *User) CheckPassword(pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
}
