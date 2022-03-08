package entity

import (
	model "github.com/AutoResto/module/recipe/entity"
)

type Menu struct {
	Id     int          `form:"id" json:"id"`
	Name   string       `form:"name" json:"name"`
	Price  float64      `form:"price" json:"price"`
	Recipe model.Recipe `form:"recipe" json:"recipe"`
}

type MenuResponse struct {
	Message string `form:"message" json:"message"`
	Data    []Menu `form:"data" json:"data"`
}