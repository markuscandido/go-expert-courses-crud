package graphql

//go:generate go run github.com/99designs/gqlgen generate
import "github.com/markuscandido/go-expert-courses-crud/internal/application/core/port"

type MutationsResolver struct {
	CreateCategoryUseCase port.ICreateCategoryUseCase
	CreateCourseUseCase   port.ICreateCourseUseCase
}

type QueriesResolver struct {
	ListCategoriesUseCase   port.IListCategoriesUseCase
	ListCoursesUseCase      port.IListCoursesUseCase
	ListCoursesByCategoryId port.IListCoursesByCategoryIdUseCase
	GetCategoryByCourseId   port.IGetCategoryByCourseIdUseCase
}

type Resolver struct {
	Mutations MutationsResolver
	Queries   QueriesResolver
}
