package entity

import (
	entityMaterial "github.com/AutoResto/module/material/entity"
)

type RecipeDetail struct {
	Id       int                     `form:"id" json:"id"`
	Material entityMaterial.Material `form:"idMaterial" json:"idMaterial"`
	Recipe   Recipe                  `form:"idRecipe" json:"idRecipe"`
	Quantity int                     `form:"quality" json:"quality"`
	Unit     string                  `form:"unit" json:"unit"`
}

type RecipeDetailResponse struct {
	Message string         `form:"message" json:"message"`
	Data    []RecipeDetail `form:"data" json:"data"`
}

// func (r *RecipeDetail) GetQuantity() int {
// 	return r.quantity
// }

// func (r *RecipeDetail) SetQuantity(quantity int) {
// 	r.quantity = quantity
// }
