package entity

import (
	MN "github.com/AutoResto/domain/menu/entity"
)

type Recipe struct {
	Id          int     `form:"id" json:"id"`
	Description string  `form:"description" json:"description"`
	Menu        MN.Menu `form : "menu" json : "menu"`
}

type RecipeResponse struct {
	Message string   `form:"message" json:"message"`
	Data    []Recipe `form:"data" json:"data"`
}
