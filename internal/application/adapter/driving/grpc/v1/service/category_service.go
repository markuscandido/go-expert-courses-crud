package service

import (
	"context"

	"github.com/markuscandido/go-expert-courses-crud/api/grpc/course/v1/pb"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/dto"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/port"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CreateCategoryUseCase port.ICreateCategoryUseCase
}

func NewCategoryService(createCategoryUseCase port.ICreateCategoryUseCase) *CategoryService {
	return &CategoryService{
		CreateCategoryUseCase: createCategoryUseCase,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	if in.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}

	createCategoryInputDTO := dto.CreateCategoryInputDTO{
		Name:        in.Name,
		Description: in.Description,
	}

	createCategoryOutputDTO, err := c.CreateCategoryUseCase.Execute(ctx, createCategoryInputDTO)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.CategoryResponse{
		Category: &pb.Category{
			Id:          createCategoryOutputDTO.ID,
			Name:        createCategoryOutputDTO.Name,
			Description: createCategoryOutputDTO.Description,
		},
	}

	return categoryResponse, nil
}
