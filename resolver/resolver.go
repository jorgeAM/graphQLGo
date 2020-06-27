package resolver

import "github.com/jorgeAM/basicGraphql/repository"

// Resolver handles dependecies
type Resolver struct {
	*repository.Layer
}
