package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/cors"

	"github.com/dmitry-drozdov/gqlgen/_examples/chat"
	"github.com/dmitry-drozdov/gqlgen/graphql/handler"
	"github.com/dmitry-drozdov/gqlgen/graphql/handler/extension"
	"github.com/dmitry-drozdov/gqlgen/graphql/handler/transport"
	"github.com/dmitry-drozdov/gqlgen/graphql/playground"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	srv := handler.New(chat.NewExecutableSchema(chat.New()))

	srv.AddTransport(transport.SSE{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.Use(extension.Introspection{})

	http.Handle("/", playground.Handler("Todo", "/query"))
	http.Handle("/query", c.Handler(srv))

	log.Fatal(http.ListenAndServe(":8085", nil))
}
