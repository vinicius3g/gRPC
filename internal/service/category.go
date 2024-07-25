package service

import (
	"context"

	"github.com/vinicius3g/gRPC.git/internal/database"
	"github.com/vinicius3g/gRPC.git/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(CategoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: CategoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)

	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: *category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}
