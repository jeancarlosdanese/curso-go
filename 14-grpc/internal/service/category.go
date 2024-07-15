package service

import (
	"context"
	"curso-go/go-grpc/internal/database"
	"curso-go/go-grpc/internal/pb"
	"io"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	resultRow, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          resultRow.ID,
		Name:        resultRow.Name,
		Description: resultRow.Description,
	}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	resultRows, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categories []*pb.Category
	for _, row := range resultRows {
		categories = append(categories, &pb.Category{
			Id:          row.ID,
			Name:        row.Name,
			Description: row.Description,
		})
	}

	return &pb.CategoryList{Categories: categories}, nil
}

func (c CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.Category, error) {
	resultRow, err := c.CategoryDB.FindByID(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          resultRow.ID,
		Name:        resultRow.Name,
		Description: resultRow.Description,
	}, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}

	// Ler continuamente do stream
	for {
		createCategory, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err // Handle errors appropriately
		}

		// Supondo que você cria e salva a categoria aqui
		resultRow, err := c.CategoryDB.Create(createCategory.Name, createCategory.Description)
		if err != nil {
			return err
		}

		// Adiciona a categoria criada à lista de categorias
		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          resultRow.ID,
			Name:        resultRow.Name,
			Description: resultRow.Description,
		})
	}
}

func (c *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	for {
		createCategory, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// Supondo que você cria e salva a categoria aqui
		resultRow, err := c.CategoryDB.Create(createCategory.Name, createCategory.Description)
		if err != nil {
			return err
		}

		// Envia a categoria criada de volta ao cliente
		err = stream.Send(&pb.Category{
			Id:          resultRow.ID,
			Name:        resultRow.Name,
			Description: resultRow.Description,
		})
		if err != nil {
			return err
		}
	}
}
