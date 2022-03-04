package recipe

import (
	"github.com/AutoResto/backend/module/material"
)

type RecipeDetail struct {
	material material.Material `form:"material" json:"material"`
	quantity int               `form:"quality" json:"quality"`
}

type RecipeDetailResponse struct {
	Message string         `form:"message" json:"message"`
	Data    []RecipeDetail `form:"data" json:"data"`
}

func (r *RecipeDetail) GetQuantity() int {
	return r.quantity
}

func (r *RecipeDetail) SetQuantity(quantity int) {
	r.quantity = quantity
}
