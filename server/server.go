package server

import (
	"context"
	"errors"
	"fmt"
	"helloserver/backend"
	"helloserver/config"
	"helloserver/graph"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const gqlpath = "/graphql"

func Start(ctx context.Context, cfg config.Config) error {
	b := backend.NewBackend()
	resolver := graph.NewResolver(b)
	routes := http.NewServeMux()
	routes.Handle(gqlpath, handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver})))
	routes.Handle("/", playground.ApolloSandboxHandler("helloserver GraphQL Playground", gqlpath,
		playground.WithApolloSandboxEndpointIsEditable(false),
	))

	server := http.Server{
		Addr:        fmt.Sprintf("localhost:%d", cfg.Port),
		Handler:     routes,
		IdleTimeout: time.Minute,
	}

	go func() {
		log.Println("starting server")
		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down server")
	shutdownctx, cancel := context.WithTimeout(context.Background(), cfg.ServerShutdownTimeout)
	defer cancel()
	err := server.Shutdown(shutdownctx)
	if err != nil {
		log.Printf("error shutting down server: %v", err)
	}

	log.Println("server stopped")
	return nil
}
