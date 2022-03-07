package entity

import (
	"github.com/AutoResto/module/material/entity"
)

type RecipeDetail struct {
	Id 	int  `form : "id" json : "id"`
	Unit string `form : "unit" json : ""unit`
	Material entity.Material `form:"material" json:"material"`
	Recipe Recipe `form : "recipe" json : "recipe"`
	Quantity int             `form:"quality" json:"quality"`
}

type RecipeDetailResponse struct {
	Message string         `form:"message" json:"message"`
	Data    []RecipeDetail `form:"data" json:"data"`
}

