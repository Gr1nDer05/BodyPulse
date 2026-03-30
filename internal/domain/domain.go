package domain

import (
	"github.com/gofrs/uuid"
)

type MealType string

type Product struct {
	UserID    int64
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

type User struct {
	UserID    int64
	Username  string
	FirstName string
}

const (
	Breakfast MealType = "breakfast"
	Lunch     MealType = "lunch"
	Dinner    MealType = "dinner"
	Snack     MealType = "snack"
)

func NewUser(UserID int64, Username string, FirstName string) (User, error) {
	if err := ValidateUserID(UserID); err != nil {
		return User{}, nil
	}

	return User{
		UserID:    UserID,
		Username:  Username,
		FirstName: FirstName,
	}, nil
}

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

func NewProduct(UserID int64, ProductID uuid.UUID, Name string, Calories float64, Proteins float64, Fats float64, Carbs float64) (Product, error) {
	if err := ValidateProduct(UserID, ProductID, Name, Calories, Proteins, Fats, Carbs); err != nil {
		return Product{}, err
	}
	return Product{
		UserID:    UserID,
		ProductID: ProductID,
		Name:      Name,
		Calories:  Calories,
		Proteins:  Proteins,
		Fats:      Fats,
		Carbs:     Carbs,
	}, nil
}
