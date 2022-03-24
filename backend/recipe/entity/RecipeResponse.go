package entity

type RecipeResponse struct {
	Message string   `form:"message" json:"message"`
	Data    []Recipe `form:"data" json:"data"`
}

type RecipeDetailResponse struct {
	Message string         `form:"message" json:"message"`
	Data    []RecipeDetail `form:"data" json:"data"`
}
