package main

import (
	"log"
	"net/http"
	"os"
	

  
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	
	"github.com/baadjis/transferservice/graph"
	"github.com/baadjis/transferservice/graph/generated"
	"github.com/baadjis/transferservice/grah/database"
	
	
)

const defaultPort = "8080"

func main(){

	database.initDB()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: database.db,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
