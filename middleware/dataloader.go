package middleware

import (
	"context"
	"net/http"

	"github.com/jorgeAM/basicGraphql/dataloader"
	"github.com/jorgeAM/basicGraphql/repository"
	"github.com/jorgeAM/basicGraphql/utils"
)

// Dataloader save loader in context
func Dataloader(repository *repository.Layer, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loader := dataloader.NewLoader(repository)
		ctx := context.WithValue(r.Context(), utils.Loaders, loader)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
