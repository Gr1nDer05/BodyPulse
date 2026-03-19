package service

import (
	"context"
	"log/slog"

	"github.com/Gr1nDer05/BodyPulse/internal/domain"
	"github.com/Gr1nDer05/BodyPulse/internal/dto"
	"github.com/Gr1nDer05/BodyPulse/internal/repository"
	"github.com/gofrs/uuid"
)

type ProductService interface {
	CreateUserProduct(ctx context.Context, req dto.CreateUserProductRequest) (domain.Product, error)
	FindProductByName(ctx context.Context, req dto.SearchProductRequest) ([]domain.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
	log               *slog.Logger
}

func NewProductService(productRepository repository.ProductRepository, log *slog.Logger) ProductService {
	return &productService{
		productRepository: productRepository,
		log:               log,
	}
}

func (s *productService) CreateUserProduct(ctx context.Context, req dto.CreateUserProductRequest) (domain.Product, error) {
	const op = "product_service.createUserProduct"
	s.log.With(
		slog.String("op", op),
	)
	s.log.Info("Creating user product")

	productID, err := uuid.NewV7()
	if err != nil {
		return domain.Product{}, err
	}

	product, err := domain.NewProduct(productID, req.Name, req.Calories, req.Proteins, req.Fats, req.Carbs)
	if err != nil {
		return domain.Product{}, err
	}
	err = s.productRepository.CreateUserProduct(ctx, product)
	if err != nil {
		s.log.Error("Failed to create user product")
		return domain.Product{}, err
	}
	return product, nil
}

func (s *productService) FindProductByName(ctx context.Context, req dto.SearchProductRequest) ([]domain.Product, error) {
	const op = "product_service.findProductByName"
	s.log.With(
		slog.String("op", op),
	)
	s.log.Info("Finding product by name")

	products, err := s.productRepository.FindProductByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return products, nil
}
