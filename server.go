package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog/log"

	"gql_jobportal/database"
	"gql_jobportal/graph"
	"gql_jobportal/repository"
	"gql_jobportal/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {

	start, err := startAppilcation()
	if err != nil {
		log.Info().Err(err).Msg("could not startapp")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Service: start}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal().Err(http.ListenAndServe(":"+port, nil))
}
func startAppilcation() (service.UserService, error) {

	db, err := database.Open()
	if err != nil {
		return &service.Service{}, fmt.Errorf("connecting to database %w", err)
	}
	pg, err := db.DB()
	if err != nil {
		return &service.Service{}, fmt.Errorf("failed to get database instance: %w ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return &service.Service{}, fmt.Errorf("database is not connected: %w ", err)
	}
	repo, err := repository.NewRepository(db)
	if err != nil {
		return &service.Service{}, fmt.Errorf("could not initialize repo layer: %w ", err)
	}
	svc, err := service.NewService(repo)
	if err != nil {
		return &service.Service{}, fmt.Errorf("could not initialize service layer: %w ", err)
	}
	return svc, nil
}
