package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/fizx/gqlgen/example/starwars"
	"github.com/fizx/gqlgen/example/starwars/generated"
	"github.com/fizx/gqlgen/graphql"
	"github.com/fizx/gqlgen/graphql/handler"
	"github.com/fizx/gqlgen/graphql/playground"
)

func main() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(starwars.NewResolver()))
	srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
		rc := graphql.GetFieldContext(ctx)
		fmt.Println("Entered", rc.Object, rc.Field.Name)
		res, err = next(ctx)
		fmt.Println("Left", rc.Object, rc.Field.Name, "=>", res, err)
		return res, err
	})

	http.Handle("/", playground.Handler("Starwars", "/query"))
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
