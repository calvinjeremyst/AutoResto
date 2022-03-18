package entity

import (
	MN "github.com/AutoResto/menu/entity"
)

type Recipe struct {
	Id          int     `form:"id" json:"id"`
	Description string  `form:"description" json:"description"`
	Menu        MN.Menu `form : "menu" json : "menu"`
}

