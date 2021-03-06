package main

import (
	"log"
	"net/http"

	"github.com/fizx/gqlgen/example/selection"
	"github.com/fizx/gqlgen/graphql/handler"
	"github.com/fizx/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Selection Demo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(selection.NewExecutableSchema(selection.Config{Resolvers: &selection.Resolver{}})))
	log.Fatal(http.ListenAndServe(":8086", nil))
}
