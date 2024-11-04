package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/swiftee19/GiziNusa-BackEnd/database/postgresql"
	"github.com/swiftee19/GiziNusa-BackEnd/entities"
	"github.com/swiftee19/GiziNusa-BackEnd/graph"
)

func main() {
	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

	postgresql.InitDB()

	err := postgresql.DB.AutoMigrate(&entities.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration successful!")

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	port := os.Getenv("GRAPHQLPORT")
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
