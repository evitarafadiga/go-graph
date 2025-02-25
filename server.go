package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/evitarafadiga/go-graph/config"
	"github.com/evitarafadiga/go-graph/db"
	"github.com/evitarafadiga/go-graph/generated"
	"github.com/evitarafadiga/go-graph/resolvers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {

	env := config.GetEnv()
	port := env.Port
	if port == "" {
		port = config.DEFAULT_PORT
	}

	db, err := db.New(env.DBName)
	if err != nil{
		log.Fatal(err)
	}

	var router *chi.Mux = chi.NewRouter()
	router.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins: []string{"*"},
				AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
				AllowedHeaders: []string{"*"},
				AllowCredentials: true,
			},
		),
	)

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{DB: db}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
