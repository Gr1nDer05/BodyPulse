package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Gr1nDer05/BodyPulse/internal/domain"
	"github.com/Gr1nDer05/BodyPulse/internal/dto"
	"github.com/Gr1nDer05/BodyPulse/internal/repository"
	"github.com/gofrs/uuid"
)

type ProductService interface {
	CreateUserProduct(ctx context.Context, req dto.CreateUserProductRequest) (domain.Product, error)
	FindProductByName(ctx context.Context, req dto.SearchProductRequest) ([]domain.Product, error)
	ListFoodsByUser(ctx context.Context) ([]domain.Product, error)
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
	logger := s.log.With(
		slog.String("op", op),
	)
	logger.Info("Creating user product")

	productID, err := uuid.NewV7()
	if err != nil {
		return domain.Product{}, err
	}

	value := ctx.Value(userIDKey)
	userID, ok := value.(int64)
	if !ok {
		s.log.Error("failed to get userID from context", "value", value)
		return domain.Product{}, errors.New("userID not found in context")
	}

	product, err := domain.NewProduct(userID, productID, req.Name, req.Calories, req.Proteins, req.Fats, req.Carbs)
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
	logger := s.log.With(
		slog.String("op", op),
	)
	logger.Info("Finding product by name")

	products, err := s.productRepository.FindProductByName(ctx, req.Name)
	if err != nil {
		s.log.Error("Failed to find product by name")
		return nil, err
	}
	return products, nil
}

func (s *productService) ListFoodsByUser(ctx context.Context) ([]domain.Product, error) {
	const op = "product_service.listOfUsersProducts"
	logger := s.log.With(
		slog.String("op", op),
	)
	logger.Info("Getting user's products")

	value := ctx.Value(userIDKey)
	userID, ok := value.(int64)
	if !ok {
		s.log.Error("failed to get userID from context", "value", value)
		return nil, errors.New("userID not found in context")
	}

	products, err := s.productRepository.ListFoodsByUser(ctx, userID)
	if err != nil {
		s.log.Error("Failed to get user's products")
		return nil, err
	}
	return products, nil
}
