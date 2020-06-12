package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jorgeAM/basicGraphql/utils"
)

var jwtKey string = os.Getenv("JWT_KEY")

// Authentication read authorization header an save user in context
func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			next.ServeHTTP(w, r)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		clain, ok := token.Claims.(*jwt.StandardClaims)

		if !ok || clain.Valid() != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), utils.UserID, clain.Id)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
