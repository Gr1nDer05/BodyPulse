package dto

type InitData struct {
	UserID    int64  `json:"user_id" binding:"required"`
	Username  string `json:"username" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	AuthDate  int64  `json:"auth_date" binding:"required"`
	Hash      string `json:"hash" binding:"required"`
}

type CreateUserProductRequest struct {
	Name     string  `json:"name" binding:"required"`
	Calories float64 `json:"calories" binding:"required"`
	Proteins float64 `json:"proteins" binding:"required"`
	Fats     float64 `json:"fats" binding:"required"`
	Carbs    float64 `json:"carbs" binding:"required"`
}

type SearchProductRequest struct {
	Name string `json:"name" binding:"required"`
}
