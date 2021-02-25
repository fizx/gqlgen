package main

import (
	"log"
	"net/http"

	todo "github.com/fizx/gqlgen/example/config"
	"github.com/fizx/gqlgen/graphql/handler"
	"github.com/fizx/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Todo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(
		todo.NewExecutableSchema(todo.New()),
	))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
