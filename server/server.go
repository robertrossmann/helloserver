package server

import (
	"errors"
	"fmt"
	"helloserver/backend"
	"helloserver/config"
	"helloserver/graph"
	"log"
	"net"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const gqlpath = "/graphql"

func Start(cfg config.Config) error {
	b := backend.NewBackend()
	resolver := graph.NewResolver(b)
	schema := graph.NewExecutableSchema(graph.Config{Resolvers: resolver})
	server := handler.NewDefaultServer(schema)

	http.Handle("/", playground.ApolloSandboxHandler("helloserver GraphQL Playground", gqlpath,
		playground.WithApolloSandboxEndpointIsEditable(false),
	))
	http.Handle(gqlpath, server)

	log.Println("starting server")
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.Port))
	log.Printf("listening on localhost:%d", cfg.Port)

	if err != nil {
		return fmt.Errorf("TCP listen: %w", err)
	}

	if err = http.Serve(listener, nil); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	log.Println("server stopped")
	return nil
}
