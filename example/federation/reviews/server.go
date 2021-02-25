//go:generate go run ../../../testdata/gqlgen.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fizx/gqlgen/example/federation/reviews/graph"
	"github.com/fizx/gqlgen/example/federation/reviews/graph/generated"
	"github.com/fizx/gqlgen/graphql/handler"
	"github.com/fizx/gqlgen/graphql/handler/debug"
	"github.com/fizx/gqlgen/graphql/playground"
)

const defaultPort = "4003"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.Use(&debug.Tracer{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
