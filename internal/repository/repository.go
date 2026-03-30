package repository

import (
	"context"

	"github.com/Gr1nDer05/BodyPulse/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repository {
	return &repository{
		db: db,
	}
}

type ProductRepository interface {
	CreateUserProduct(ctx context.Context, product domain.Product) error
	FindProductByName(cxt context.Context, name string) ([]domain.Product, error)
	ListFoodsByUser(ctx context.Context, userID int64) ([]domain.Product, error)
}

type FoodEntryRepository interface {
}
