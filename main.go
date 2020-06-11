package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/basicGraphql/generated"
	"github.com/jorgeAM/basicGraphql/resolver"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	cfg := generated.Config{
		Resolvers: &resolver.Resolver{},
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
