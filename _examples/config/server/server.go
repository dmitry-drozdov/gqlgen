package main

import (
	"log"
	"net/http"

	todo "github.com/dmitry-drozdov/gqlgen/_examples/config"
	"github.com/dmitry-drozdov/gqlgen/graphql/handler"
	"github.com/dmitry-drozdov/gqlgen/graphql/handler/transport"
	"github.com/dmitry-drozdov/gqlgen/graphql/playground"
)

func main() {
	srv := handler.New(
		todo.NewExecutableSchema(todo.New()),
	)
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	http.Handle("/", playground.Handler("Todo", "/query"))
	http.Handle("/query", srv)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
