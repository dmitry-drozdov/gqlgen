package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dmitry-drozdov/gqlgen/_examples/starwars"
	"github.com/dmitry-drozdov/gqlgen/_examples/starwars/generated"
	"github.com/dmitry-drozdov/gqlgen/graphql"
	"github.com/dmitry-drozdov/gqlgen/graphql/handler"
	"github.com/dmitry-drozdov/gqlgen/graphql/handler/transport"
	"github.com/dmitry-drozdov/gqlgen/graphql/playground"
)

func main() {
	srv := handler.New(generated.NewExecutableSchema(starwars.NewResolver()))
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res any, err error) {
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
