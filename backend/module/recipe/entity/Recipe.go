package entity

type Recipe struct {
	Id          int    `form:"id" json:"id"`
	Description string `form:"description" json:"description"`
}

type RecipeResponse struct {
	Message string   `form:"message" json:"message"`
	Data    []Recipe `form:"data" json:"data"`
}
