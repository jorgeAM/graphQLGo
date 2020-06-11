package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/basicGraphql/db"
	"github.com/jorgeAM/basicGraphql/generated"
	userrepository "github.com/jorgeAM/basicGraphql/repositories/user"
	"github.com/jorgeAM/basicGraphql/resolver"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	dbEngine := os.Getenv("DB_ENGINE")
	dbURL := os.Getenv("DB_URL")

	dbHandler, err := db.NewPersistenceLayer(db.TYPE(dbEngine), dbURL)

	if err != nil {
		log.Fatalf("Something get wrong to connect to %s: %v", dbEngine, err)
	}

	userRep, err := userrepository.NewUserRepository(db.TYPE(dbEngine), dbHandler)

	if err != nil {
		log.Fatalf("Something get wrong to initialize user repository to %s: %v", dbEngine, err)
	}

	cfg := generated.Config{
		Resolvers: &resolver.Resolver{UserResolver: userRep},
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
