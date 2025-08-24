package graphql

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/markuscandido/go-expert-courses-crud/internal/application"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/adapter/driven/config"
)

func StartGraphQLServer(cfg *config.Config, useCases *application.UseCases) {
	resolver := &Resolver{
		Mutations: MutationsResolver{
			CreateCategoryUseCase: useCases.CreateCategoryUseCase,
			CreateCourseUseCase:   useCases.CreateCourseUseCase,
		},
		Queries: QueriesResolver{
			ListCategoriesUseCase:   useCases.ListCategoriesUseCase,
			ListCoursesUseCase:      useCases.ListCoursesUseCase,
			GetCategoryByCourseId:   useCases.GetCategoryByCourseIdUseCase,
			ListCoursesByCategoryId: useCases.ListCoursesByCategoryIdUseCase,
		},
	}

	srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("GraphQL playground available at http://localhost:%s/", cfg.GraphQLPort)
	log.Printf("Starting GraphQL server on port %s...", cfg.GraphQLPort)
	if err := http.ListenAndServe(":"+cfg.GraphQLPort, nil); err != nil {
		log.Fatalf("FATAL: could not start GraphQL server: %v", err)
	}
}
