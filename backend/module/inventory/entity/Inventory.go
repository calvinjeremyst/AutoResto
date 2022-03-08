package entity

import (
	model "github.com/AutoResto/module/location/entity"
)

type Inventory struct {
	Id       int            `form:"id" json:"id"`
	Capacity int            `form:"capacity" json:"capacity"`
	Location model.Location `form:"location" json:"location"`
}

type InventoryResponse struct {
	Message string      `form:"message" json:"message"`
	Data    []Inventory `form:"data" json:"data"`
}
