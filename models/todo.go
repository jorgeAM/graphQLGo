package models

import "time"

// Todo model
type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" pg:"type:varchar(100)"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt" pg:"default:now()"`
	UpdatedAt   time.Time `json:"updatedAt" pg:"default:now()"`
	UserID      int       `json:"user" pg:"fk:user_id"`
}
