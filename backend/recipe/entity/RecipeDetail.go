package entity

import (
	model "github.com/AutoResto/material/entity"
)

type RecipeDetail struct {
	Id       int            `form:"id" json:"id"`
	Material model.Material `form:"idMaterial" json:"idMaterial"`
	Recipe   Recipe         `form:"idRecipe" json:"idRecipe"`
	Quantity int            `form:"quality" json:"quality"`
	Unit     string         `form:"unit" json:"unit"`
}
