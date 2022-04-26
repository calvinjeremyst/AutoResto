package entity

import (
	model "github.com/AutoResto/material/entity"
)

type RecipeDetail struct {
	Id       int            `form:"id" json:"id"`
	Material model.Material `form:"material" json:"material"`
	Recipe   Recipe         `form:"recipe" json:"recipe"`
	Quantity int            `form:"quantity" json:"quantity"`
	Unit     string         `form:"unit" json:"unit"`
}
