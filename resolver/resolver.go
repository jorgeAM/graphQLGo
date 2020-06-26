package resolver

import (
	todorepository "github.com/jorgeAM/basicGraphql/repositories/todo"
	userresolver "github.com/jorgeAM/basicGraphql/repositories/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver handles dependecies
type Resolver struct {
	UserResolver userresolver.Handler
	TodoResolver todorepository.Handler
}
