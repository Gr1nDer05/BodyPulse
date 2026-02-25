package models

type Product struct {
	Name     string  `json:"name"`
	Calories float64 `json:"calories"`
	Proteins float64 `json:"proteins"`
	Fats     float64 `json:"fats"`
	Carbs    float64 `json:"carbs"`
}

type ProductEntry struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}
