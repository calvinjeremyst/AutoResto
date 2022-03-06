package entity

import (
	"github.com/AutoResto/module/material/entity"
)

type RecipeDetail struct {
	Material entity.Material `form:"material" json:"material"`
	Quantity int             `form:"quality" json:"quality"`
}

type RecipeDetailResponse struct {
	Message string         `form:"message" json:"message"`
	Data    []RecipeDetail `form:"data" json:"data"`
}
