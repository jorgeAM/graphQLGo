package models

import "time"

// User model
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" pg:"type:varchar(100)"`
	Email     string    `json:"email" pg:",unique,type:varchar(100)"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt" pg:"default:now()"`
	UpdatedAt time.Time `json:"updatedAt" pg:"default:now()"`
}
