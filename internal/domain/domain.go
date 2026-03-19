package domain

import (
	"github.com/gofrs/uuid"
)

type MealType string

type Product struct {
	ProductID uuid.UUID
	Name      string
	Calories  float64
	Proteins  float64
	Fats      float64
	Carbs     float64
}

type FoodEntry struct {
	UserID    int64
	ProductID uuid.UUID
	Weight    int64
	MealType  MealType
}

const (
	Breakfast MealType = "breakfast"
	Lunch     MealType = "lunch"
	Dinner    MealType = "dinner"
	Snack     MealType = "snack"
)

func NewFoodEntry(UserID int64, ProductID uuid.UUID, Weight int64, MealType MealType) (FoodEntry, error) {
	if err := ValidateFoodEntry(UserID, ProductID, Weight, MealType); err != nil {
		return FoodEntry{}, err
	}

	return FoodEntry{
		UserID:    UserID,
		ProductID: ProductID,
		Weight:    Weight,
		MealType:  MealType,
	}, nil
}

func NewProduct(ProductID uuid.UUID, Name string, Calories float64, Proteins float64, Fats float64, Carbs float64) (Product, error) {
	if err := ValidateProduct(ProductID, Name, Calories, Proteins, Fats, Carbs); err != nil {
		return Product{}, err
	}
	return Product{
		ProductID: ProductID,
		Name:      Name,
		Calories:  Calories,
		Proteins:  Proteins,
		Fats:      Fats,
		Carbs:     Carbs,
	}, nil
}
