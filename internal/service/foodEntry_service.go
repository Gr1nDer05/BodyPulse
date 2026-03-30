package service

import (
	"context"
	"log/slog"

	"github.com/Gr1nDer05/BodyPulse/internal/repository"
)

type FoodEntryService interface {
	AddFood(ctx context.Context) error
}

type foodEntryService struct {
	foodEntryRepository repository.FoodEntryRepository
	log                 *slog.Logger
}

func NewFoodEntryService(foodEntryRepository repository.FoodEntryRepository, log *slog.Logger) FoodEntryService {
	return &foodEntryService{
		foodEntryRepository: foodEntryRepository,
		log:                 log,
	}
}

func (s *foodEntryService) AddFood(ctx context.Context) error {
	return nil
}
