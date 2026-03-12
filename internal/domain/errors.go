package domain

import (
	"errors"

	"github.com/gofrs/uuid"
)

var (
	ErrInvalidUserID    = errors.New("invalid user id")
	ErrInvalidProductID = errors.New("invalid product id")
	ErrInvalidWeight    = errors.New("invalid weight")
	ErrInvalidMealType  = errors.New("invalid meal type")
	ErrInvalidProteins  = errors.New("invalid proteins")
	ErrInvalidFats      = errors.New("invalid fats")
	ErrInvalidCarbs     = errors.New("invalid carbs")
	ErrInvalidCalories  = errors.New("invalid calories")
	ErrInvalidName      = errors.New("invalid product name")
)

func (m MealType) isValid() bool {
	switch m {
	case Breakfast, Lunch, Dinner, Snack:
		return true
	default:
		return false
	}
}

func ValidateFoodEntry(UserID int64, ProductID uuid.UUID, Weight int64, MealType MealType) error {
	if UserID <= 0 {
		return ErrInvalidUserID
	}
	if ProductID == uuid.Nil {
		return ErrInvalidProductID
	}
	if Weight <= 0 {
		return ErrInvalidWeight
	}
	if !MealType.isValid() {
		return ErrInvalidMealType
	}
	return nil
}

func ValidateProduct(Name string, ProductID uuid.UUID, Calories float64, Proteins float64, Fats float64, Carbs float64) error {
	if ProductID == uuid.Nil {
		return ErrInvalidProductID
	}
	if Name == "" {
		return ErrInvalidName
	}
	if Calories <= 0 {
		return ErrInvalidCalories
	}
	if Proteins < 0 {
		return ErrInvalidProteins
	}
	if Fats < 0 {
		return ErrInvalidFats
	}
	if Carbs < 0 {
		return ErrInvalidCarbs
	}
	return nil
}
