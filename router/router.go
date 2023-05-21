package router

import (
	"github.com/ahmedkhalaf1996/GRQGoTaskx/graph"
	"github.com/ahmedkhalaf1996/GRQGoTaskx/graph/generated"

	customMiddleware "github.com/ahmedkhalaf1996/GRQGoTaskx/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

func CallRouter(db *gorm.DB) *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	router.Use(middleware.RequestID)
	router.Use(customMiddleware.AuthMiddleware(db))
	router.Use(graph.UserDataloaderMiddleware(db))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	return router
}
