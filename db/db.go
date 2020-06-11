package db

import (
	"fmt"
)

// TYPE is a custom type to handle database engines
type TYPE string

// Handler implement function that all database engine should implement
type Handler interface {
	Close() error
}

// database engine supported
const (
	POSTGRES TYPE = "postgres"
	MONGODB  TYPE = "mongodb"
)

// NewPersistenceLayer initialize db handler
func NewPersistenceLayer(engine TYPE, url string) (Handler, error) {
	switch engine {
	case POSTGRES:
		return NewPostgresHandler(url)
	default:
		return nil, fmt.Errorf("%s engine is not supported yet", engine)
	}
}
