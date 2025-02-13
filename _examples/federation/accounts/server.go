//go:generate go run ../../../testdata/gqlgen.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dmitry-drozdov/gqlgen/_examples/federation/accounts/graph"
	"github.com/dmitry-drozdov/gqlgen/graphql/handler"
	"github.com/dmitry-drozdov/gqlgen/graphql/handler/debug"
	"github.com/dmitry-drozdov/gqlgen/graphql/handler/extension"
	"github.com/dmitry-drozdov/gqlgen/graphql/handler/transport"
	"github.com/dmitry-drozdov/gqlgen/graphql/playground"
)

const defaultPort = "4001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(
		graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}),
	)
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{})
	srv.Use(&debug.Tracer{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
